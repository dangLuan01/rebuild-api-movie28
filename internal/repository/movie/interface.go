package movierepository

import (
	v1dto "github.com/dangLuan01/rebuild-api-movie28/internal/dto/v1"
)

type MovieRepository interface {
	FindByHot(limit int64) ([]v1dto.MovieRawDTO, error)
	FindAll(page, pageSize int64) ([]v1dto.MovieRawDTO, v1dto.Paginate, error)
	FindBySlug(slug string) (*v1dto.MovieDetailDTO, error)
	Filter(filter *v1dto.Filter, page, pageSize int64) ([]v1dto.MovieRawDTO, *v1dto.Paginate, error)
	SiteMap(types string) ([]v1dto.SiteMap, error)
}