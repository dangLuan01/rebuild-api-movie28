package v1service

import (
	"fmt"

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
	Id int64
	PageTheme int64
	PageMovie int64
	Limit int64
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

	var themes *v1dto.ThemesWithPaginateDTO
	getTheme := ts.rd.Get(fmt.Sprintf("themes:page_theme=%v:page_movie=%v:limit=%v",
				param.PageTheme,
				param.PageMovie,
				param.Limit,
	), &themes)

	if !getTheme {
		themes, err := ts.repo.FindAll(param.Id, param.PageTheme, param.PageMovie, param.Limit)
		if err != nil {

			return nil, utils.WrapError(
				string(utils.ErrCodeInternal),
				"Faile get all theme",
				err,
			)
		}
		if err := ts.rd.Set(fmt.Sprintf("themes:page_theme=%v:page_movie=%v:limit=%v",
			param.PageTheme,
			param.PageMovie,
			param.Limit,
		), &themes); err != nil {
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