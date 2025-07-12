package themerepository

import (
	v1dto "github.com/dangLuan01/rebuild-api-movie28/internal/dto/v1"
)

type ThemeRepository interface {
	FindAll(id, pageTheme, pageMovie, limit int64) (*v1dto.ThemesWithPaginateDTO, error)
}