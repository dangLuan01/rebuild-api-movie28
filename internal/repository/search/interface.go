package searchrepository

import (
	"bytes"

	"github.com/dangLuan01/rebuild-api-movie28/internal/models"
)

type SearchRepository interface {
	Search(index string, search bytes.Buffer) ([]models.Movie, error)
}