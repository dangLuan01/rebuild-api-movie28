package v1service

import (
	"github.com/dangLuan01/rebuild-api-movie28/internal/models"
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

func (ms *movieService) GetMovieHot(limit int) ([]models.Movie, error) {
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