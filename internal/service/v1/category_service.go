package v1service

import (
	"log"

	"github.com/dangLuan01/rebuild-api-movie28/internal/models"
	categoryrepository "github.com/dangLuan01/rebuild-api-movie28/internal/repository/category"
	"github.com/dangLuan01/rebuild-api-movie28/internal/utils"
	"github.com/dangLuan01/rebuild-api-movie28/pkg/cache"
	"github.com/redis/go-redis/v9"
)

type categoryService struct {
	repo categoryrepository.CategoryRepository
	cache *cache.RedisCacheService
}

func NewCategoryService(repo categoryrepository.CategoryRepository, redisClient *redis.Client) CategoryService {
	return &categoryService{
		repo: repo,
		cache: cache.NewRedisCacheService(redisClient),
	}
}

func (cs *categoryService) GetAllCategory() ([]models.Category, error) {
	var categories []models.Category
	cacheCategories := cs.cache.Get("categories", &categories)

	if cacheCategories != redis.Nil && cacheCategories == nil {
		return categories, nil
	}
	categories, err := cs.repo.FindAll()
	if err != nil {
		return nil, utils.WrapError(
			string(utils.ErrCodeInternal),
			"Faile get all category",
			err,
		)
	}
	if err := cs.cache.Set("categories", categories, 0); err != nil {
		log.Printf("❌ Failed category set to cache:%v", err)
	}

	return categories, nil
}