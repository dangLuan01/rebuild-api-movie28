package movierepository

import (
	v1dto "github.com/dangLuan01/rebuild-api-movie28/internal/dto/v1"
)

type MovieRepository interface {
	FindByHot(limit int) ([]v1dto.MovieRawDTO, error)
	FindAll(page, pageSize int) ([]v1dto.MovieRawDTO, v1dto.Paginate, error)
	FindBySlug(slug string) (*v1dto.MovieDetailDTO, error)
}