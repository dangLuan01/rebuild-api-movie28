package v1service

import (
	"github.com/dangLuan01/rebuild-api-movie28/internal/models"
	//"github.com/dangLuan01/rebuild-api-movie28/internal/repository"
	genrerepository "github.com/dangLuan01/rebuild-api-movie28/internal/repository/genre"
	"github.com/dangLuan01/rebuild-api-movie28/internal/repository/redis"
	"github.com/dangLuan01/rebuild-api-movie28/internal/utils"
)

type genreService struct {
	repo genrerepository.GenreRepository
	rd   redis.RedisRepository
}

func NewGenreService(repo genrerepository.GenreRepository, rd redis.RedisRepository) GenreService {
	return &genreService{
		repo: repo,
		rd:   rd,
	}
}

func (gs *genreService) GetAllGenres() ([]models.Genre, error) {

	var genres []models.Genre
	genreCache := gs.rd.Get("genres", &genres)
	if !genreCache {
		genres, err := gs.repo.FindAll()
		if err != nil {
			
			return nil, utils.WrapError(
				string(utils.ErrCodeInternal), 
				"Faile fetch users.", 
				err,
			)
		}
		err = gs.rd.Set("genres", genres)
		if err != nil {
			return nil, utils.WrapError(
				string(utils.ErrCodeInternal), 
				"Faile set cache genre", 
				err,
			)
		}
		return genres, nil
	}
	
	return genres, nil
}