package v1service

import (
	"fmt"
	"log"

	v1dto "github.com/dangLuan01/rebuild-api-movie28/internal/dto/v1"
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
func (gs *genreService)GetGenreBySlug(slug string, page, pageSize int64) ([]v1dto.MovieRawDTO, models.Genre, v1dto.Paginate, error) {
	if page == 0  {
		page = 1
	}
	if pageSize == 0 {
		pageSize = 20
	}
	var GenreBySlug struct {
		Genre models.Genre `json:"genre"`
		Movies []v1dto.MovieRawDTO `json:"movies"`
		Paginate v1dto.Paginate `json:"paginate"`
	}
	key := fmt.Sprintf("genres:slug=%s:page=%d:pageSize=%d", slug, page, pageSize)
	cacheGenreBySlug := gs.cache.Get(key, &GenreBySlug)
	if cacheGenreBySlug != redis.Nil && cacheGenreBySlug == nil {
		return GenreBySlug.Movies, GenreBySlug.Genre, GenreBySlug.Paginate, nil
	}

	movie, genre, paginate, err := gs.repo.FindBySlug(slug, page, pageSize)
	if err != nil {
		return nil, models.Genre{}, v1dto.Paginate{}, utils.WrapError(
			string(utils.ErrCodeInternal),
			"Faile get genre with list movie",
			err,
		)
	}
	GenreBySlug = struct {
		Genre models.Genre `json:"genre"`
		Movies []v1dto.MovieRawDTO `json:"movies"`
		Paginate v1dto.Paginate `json:"paginate"`
	} {
		Genre: genre,
		Movies: movie,
		Paginate: paginate,
	}
	if err := gs.cache.Set(key, GenreBySlug, utils.RandomTimeSecond()); err != nil {
		log.Printf("❌ Failed genre by slug set cache:%v", err)
	}
	return movie, genre, paginate, nil
}