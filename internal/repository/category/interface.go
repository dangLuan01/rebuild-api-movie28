package categoryrepository

import "github.com/dangLuan01/rebuild-api-movie28/internal/models"

type CategoryRepository interface {
	FindAll() ([]models.Category, error)
}