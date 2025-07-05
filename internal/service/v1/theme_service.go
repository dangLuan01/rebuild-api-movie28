package v1service

import (
	v1dto "github.com/dangLuan01/rebuild-api-movie28/internal/dto/v1"
	"github.com/dangLuan01/rebuild-api-movie28/internal/repository/redis"
	themerepository "github.com/dangLuan01/rebuild-api-movie28/internal/repository/theme"
	"github.com/dangLuan01/rebuild-api-movie28/internal/utils"
)

type themeService struct {
	repo themerepository.ThemeRepository
	rd   redis.RedisRepository
}

type ThemeParam struct {
	Id int
	PageTheme int
	PageMovie int
	Limit int
}

func NewThemeService (repo themerepository.ThemeRepository, rd redis.RedisRepository) ThemeService {
	return &themeService{
		repo: repo,
		rd:   rd,
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

	themes, err := ts.repo.FindAll(param.Id, param.PageTheme, param.PageMovie, param.Limit)
	if err != nil {
		return nil, utils.WrapError(
			string(utils.ErrCodeInternal),
			"Faile get all theme",
			err,
		)
	}

	return themes, nil
}