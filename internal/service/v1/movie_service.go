package v1service

import (
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
	movies, err := ms.repo.FindByHot(limit)
	if err != nil {
		return nil, utils.WrapError(
			string(utils.ErrCodeInternal), 
			"Faile fetch movie hot.", 
			err,
		)
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