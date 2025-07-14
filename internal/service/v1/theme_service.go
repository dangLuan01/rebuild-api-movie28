package v1service

import (
	"fmt"
	"time"

	v1dto "github.com/dangLuan01/rebuild-api-movie28/internal/dto/v1"
	themerepository "github.com/dangLuan01/rebuild-api-movie28/internal/repository/theme"
	"github.com/dangLuan01/rebuild-api-movie28/internal/utils"
	"github.com/dangLuan01/rebuild-api-movie28/pkg/cache"
	"github.com/redis/go-redis/v9"
)

type themeService struct {
	repo themerepository.ThemeRepository
	cache   *cache.RedisCacheService
}

type ThemeParam struct {
	Id int64
	PageTheme int64
	PageMovie int64
	Limit int64
}

func NewThemeService (repo themerepository.ThemeRepository, redisClient *redis.Client) ThemeService {
	return &themeService{
		repo: repo,
		cache: cache.NewRedisCacheService(redisClient),
	}
}

func (ts *themeService) GetAllThemes (param ThemeParam) (*v1dto.ThemesWithPaginateDTO, error) {
    if param.PageTheme == 0 {
		param.PageTheme = 1
	}

	if param.PageMovie == 0 {
		param.PageMovie = 1
	}

	if param.Limit == 0 {
		param.Limit = 2
	}

	var themes *v1dto.ThemesWithPaginateDTO
	getTheme := ts.cache.Get(fmt.Sprintf("themes:page_theme=%v:page_movie=%v:limit=%v",
				param.PageTheme,
				param.PageMovie,
				param.Limit,
	), &themes)

	if getTheme != nil {
		themes, err := ts.repo.FindAll(param.Id, param.PageTheme, param.PageMovie, param.Limit)
		if err != nil {

			return nil, utils.WrapError (
				string(utils.ErrCodeInternal),
				"Faile get all theme",
				err,
			)
		}
		if err := ts.cache.Set(fmt.Sprintf("themes:page_theme=%v:page_movie=%v:limit=%v",
			param.PageTheme,
			param.PageMovie,
			param.Limit,
		), &themes, 3 * time.Minute); err != nil {
			return nil, utils.WrapError(
				string(utils.ErrCodeInternal),
				"Faile set cache theme to redis",
				err,
			)
		}

		return themes, nil
	}
	
	return themes, nil
}