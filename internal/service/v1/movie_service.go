package v1service

import (
	"fmt"
	"strings"

	v1dto "github.com/dangLuan01/rebuild-api-movie28/internal/dto/v1"
	movierepository "github.com/dangLuan01/rebuild-api-movie28/internal/repository/movie"
	"github.com/dangLuan01/rebuild-api-movie28/internal/utils"
	"github.com/dangLuan01/rebuild-api-movie28/pkg/cache"
	"github.com/redis/go-redis/v9"
)

type movieService struct {
	repo movierepository.MovieRepository
	cache   *cache.RedisCacheService
}

func NewMovieService(repo movierepository.MovieRepository, redisClient *redis.Client) MovieService {
	return &movieService{
		repo: repo,
		cache: cache.NewRedisCacheService(redisClient),
	}
}

func (ms *movieService) GetMovieHot(limit int64) ([]v1dto.MovieRawDTO, error) {
	if limit == 0 {
		limit = 10
	}
	var movies []v1dto.MovieRawDTO
	hotCache := ms.cache.Get("movies-hot", &movies)
	if hotCache != nil {
		movies, err := ms.repo.FindByHot(limit)
		if err != nil {
			return nil, utils.WrapError(
				string(utils.ErrCodeInternal), 
				"Faile fetch movie hot.", 
				err,
			)
		}

		if err := ms.cache.Set("movies-hot", movies, utils.RandomTimeSecond()); err != nil{
			return nil, utils.WrapError(
				string(utils.ErrCodeInternal),
				"Failed set cache movie hot to redis",
				err,
			)
		}
		
		return movies, nil
	}

	return movies, nil
}

func (ms *movieService) GetAllMovies(page, pageSize int64) ([]v1dto.MovieRawDTO, v1dto.Paginate, error) {
	if page == 0 {
		page = 1
	}
	if pageSize == 0 {
		pageSize = 18
	}

	movies, paginate, err := ms.repo.FindAll(page, pageSize)
	if err != nil {
		return nil, v1dto.Paginate{}, utils.WrapError(
			string(utils.ErrCodeInternal),
			"Faile fetch all movies.",
			err,
		)
	}

	return movies, v1dto.Paginate{
		Page: paginate.Page,
		PageSize: paginate.PageSize,
		TotalPages: paginate.TotalPages,
	}, nil
}

func (ms *movieService) GetMovieDetail(slug string) (*v1dto.MovieDetailDTO, error) {

	movie, err :=ms.repo.FindBySlug(slug)

	er := fmt.Sprintln(err)
	if strings.Contains(er,"Not found") {
		return nil, utils.WrapError(
			string(utils.ErrCodeNotFound),
			"Fetch movie detail not found",
			err,
		)
	}
	if err != nil {
		return nil, utils.WrapError(
			string(utils.ErrCodeInternal),
			"Fetch movie detail error",
			err,
		)
	}

	return movie, nil
}

func (ms *movieService) FilterMovie(filter *v1dto.Filter, page, pageSize int64) ([]v1dto.MovieRawDTO, v1dto.Paginate, error){
	
	var cacheFilter struct{
		MovieFilter []v1dto.MovieRawDTO `json:"movie_filter"`
		Paginate v1dto.Paginate 		`json:"paginate"`
	}

	if page == 0 {
		page = 1
	}
	if pageSize == 0 {
		pageSize = 18
	}

	if filter.Genre != nil {
		*filter.Genre = "all"
	}

	keyFilter 	:= fmt.Sprintf("movieFilter:genre=%v:year=%v:type=%v:page=%d:pageSize=%d",
					*filter.Genre, *filter.Release_date, *filter.Type, page, pageSize)

	movieFilterCache := ms.cache.Get(keyFilter, &cacheFilter) 

	if movieFilterCache == nil && cacheFilter.MovieFilter != nil {

		return cacheFilter.MovieFilter, cacheFilter.Paginate, nil
	}

	MovieFilter, paginate, err := ms.repo.Filter(filter, page, pageSize)
	er := fmt.Sprintln(err)
	if strings.Contains(er,"Not found") {
		return nil, v1dto.Paginate{}, utils.WrapError (
			string(utils.ErrCodeNotFound),
			"Fetch movie filter not found",
			err,
		)
	}
	if err != nil {
		return nil, v1dto.Paginate{}, utils.WrapError (
			string(utils.ErrCodeInternal),
			"Faile get movie filter",
			err,
		)
	}
	cacheFilter = struct {
		MovieFilter []v1dto.MovieRawDTO `json:"movie_filter"`
		Paginate v1dto.Paginate 		`json:"paginate"`
	} {
		MovieFilter: MovieFilter,
		Paginate: v1dto.Paginate {
			Page: paginate.Page,
			PageSize: paginate.PageSize,
			TotalPages: paginate.TotalPages,
		},
	}
	
	if err:= ms.cache.Set(keyFilter, cacheFilter, utils.RandomTimeSecond()); err != nil {
		return nil, v1dto.Paginate{}, utils.WrapError (
			string(utils.ErrCodeInternal),
			"Faile set cache movie filter",
			err,
		)
	}
	
	return cacheFilter.MovieFilter, cacheFilter.Paginate, nil
}