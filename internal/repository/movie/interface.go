package movierepository

import "github.com/dangLuan01/rebuild-api-movie28/internal/models"

type MovieRepository interface {
	FindByHot(limit int) ([]models.Movie, error)
}