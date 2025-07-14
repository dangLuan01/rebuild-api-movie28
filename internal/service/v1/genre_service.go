package v1service

import (
	"github.com/dangLuan01/rebuild-api-movie28/internal/models"
	genrerepository "github.com/dangLuan01/rebuild-api-movie28/internal/repository/genre"
	"github.com/dangLuan01/rebuild-api-movie28/internal/utils"
	"github.com/dangLuan01/rebuild-api-movie28/pkg/cache"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type genreService struct {
	repo genrerepository.GenreRepository
	cache *cache.RedisCacheService
}

func NewGenreService(repo genrerepository.GenreRepository, redisClient *redis.Client) GenreService {
	return &genreService{
		repo: repo,
		cache: cache.NewRedisCacheService(redisClient),
	}
}

func (gs *genreService) GetAllGenres(ctx *gin.Context) ([]models.Genre, error) {
	
	var genres []models.Genre
	genreCache := gs.cache.Get("genres", &genres)
	if genreCache != nil {
		genres, err := gs.repo.FindAll()
		if err != nil {
			
			return nil, utils.WrapError(
				string(utils.ErrCodeInternal), 
				"Faile fetch genre.", 
				err,
			)
		}
		err = gs.cache.Set("genres", genres, 0)
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
func (gs *genreService)GetGenreBySlug(slug string, page, pageSize int64) (models.GenreWithMovie, error) {
	if page == 0  {
		page = 1
	}
	if pageSize == 0 {
		pageSize = 20
	}
	genre, err := gs.repo.FindBySlug(slug, page, pageSize)
	if err != nil {
		return models.GenreWithMovie{}, utils.WrapError(
			string(utils.ErrCodeInternal),
			"Faile get genre with list movie",
			err,
		)
	}
	
	return genre, nil
}