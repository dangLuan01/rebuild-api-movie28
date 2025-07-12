package v1service

import (
	"fmt"
	"strings"

	v1dto "github.com/dangLuan01/rebuild-api-movie28/internal/dto/v1"
	movierepository "github.com/dangLuan01/rebuild-api-movie28/internal/repository/movie"
	"github.com/dangLuan01/rebuild-api-movie28/internal/repository/redis"
	"github.com/dangLuan01/rebuild-api-movie28/internal/utils"
)

type movieService struct {
	repo movierepository.MovieRepository
	rd   redis.RedisRepository
}

func NewMovieService(repo movierepository.MovieRepository, rd redis.RedisRepository) MovieService {
	return &movieService{
		repo: repo,
		rd:   rd,
	}
}

func (ms *movieService) GetMovieHot(limit int) ([]v1dto.MovieRawDTO, error) {
	if limit == 0 {
		limit = 10
	}
	var movies []v1dto.MovieRawDTO
	hotCache := ms.rd.Get("movies-hot", &movies)
	if !hotCache {
		movies, err := ms.repo.FindByHot(limit)
		if err != nil {
			return nil, utils.WrapError(
				string(utils.ErrCodeInternal), 
				"Faile fetch movie hot.", 
				err,
			)
		}

		if err := ms.rd.Set("movies-hot", movies); err != nil{
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

func (ms *movieService) GetAllMovies(page, pageSize int) ([]v1dto.MovieRawDTO, v1dto.Paginate, error) {
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

func (ms *movieService) FilterMovie(filter *v1dto.Filter, page, pageSize int) ([]v1dto.MovieRawDTO, *v1dto.Paginate, error){
<<<<<<< HEAD
	
	var (
		movieFilter []v1dto.MovieRawDTO
		paginate *v1dto.Paginate
	)

=======
>>>>>>> 18fc222ea0344fd532603d382becf752924f558b
	if page == 0 {
		page = 1
	}
	if pageSize == 0 {
		pageSize = 18
	}

<<<<<<< HEAD
	keyFilter 	:= fmt.Sprintf("movieFilter:genre=%v:year=%v:type=%v",
					*filter.Genre, *filter.Release_date, *filter.Type)
	keyPaginate := fmt.Sprintf("moviePaginate:genre=%v:year=%v:type=%v:page=%d:pageSize=%d",
					*filter.Genre, *filter.Release_date, *filter.Type, page, pageSize)

	movieFilterCache 	:= ms.rd.Get(keyFilter, movieFilter)
	moviePaginateCache 	:= ms.rd.Get(keyPaginate, &paginate)

	if !movieFilterCache || !moviePaginateCache {
		movieFilter, paginate, err := ms.repo.Filter(filter, page, pageSize)
		er := fmt.Sprintln(err)
		if strings.Contains(er,"Not found") {
			return nil, nil, utils.WrapError (
				string(utils.ErrCodeNotFound),
				"Fetch movie filter not found",
				err,
			)
		}
		if err != nil {
			return nil, nil, utils.WrapError (
				string(utils.ErrCodeInternal),
				"Faile get movie filter",
				err,
			)
		}
		
		if err:= ms.rd.Set(keyFilter, movieFilter); err != nil {
			return nil, nil, utils.WrapError (
				string(utils.ErrCodeInternal),
				"Faile set cache movie filter",
				err,
			)
		}
		
		if err:= ms.rd.Set(keyPaginate, *paginate); err != nil {
			return nil, nil, utils.WrapError (
				string(utils.ErrCodeInternal),
				"Faile set cache movie filter",
				err,
			)
		}
		
		return movieFilter, &v1dto.Paginate {
			Page: paginate.Page,
			PageSize: paginate.PageSize,
			TotalPages: paginate.TotalPages,
		}, nil
	}
	
=======
	movieFilter, paginate, err := ms.repo.Filter(filter, page, pageSize)
	er := fmt.Sprintln(err)
	if strings.Contains(er,"Not found") {
		return nil, nil, utils.WrapError(
			string(utils.ErrCodeNotFound),
			"Fetch movie filter not found",
			err,
		)
	}
	if err != nil {
		return nil, nil, utils.WrapError(
			string(utils.ErrCodeInternal),
			"Faile get movie filter",
			err,
		)
	}

>>>>>>> 18fc222ea0344fd532603d382becf752924f558b
	return movieFilter, &v1dto.Paginate{
		Page: paginate.Page,
		PageSize: paginate.PageSize,
		TotalPages: paginate.TotalPages,
	}, nil

}