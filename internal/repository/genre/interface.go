package genrerepository

import "github.com/dangLuan01/rebuild-api-movie28/internal/models"

type GenreRepository interface {
	FindAll() ([]models.Genre, error)
}