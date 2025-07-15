package genrerepository

import (
	v1dto "github.com/dangLuan01/rebuild-api-movie28/internal/dto/v1"
	"github.com/dangLuan01/rebuild-api-movie28/internal/models"
)

type GenreRepository interface {
	FindAll() ([]models.Genre, error)
	FindBySlug(slug string, page, pageSize int64) ([]v1dto.MovieRawDTO, models.Genre, v1dto.Paginate, error)
}