package genrerepository

import (
	"fmt"

	"github.com/dangLuan01/rebuild-api-movie28/internal/models"
	"github.com/doug-martin/goqu/v9"
)

type SqlGenreRepository struct {
	genres []models.Genre
	db    *goqu.Database
}

func NewSqlGenreRepository(DB *goqu.Database) GenreRepository {
	return &SqlGenreRepository{
		genres: make([]models.Genre, 0),
		db:    DB,
	}
}

func (ur *SqlGenreRepository) FindAll() ([]models.Genre, error) {
	
	ds := ur.db.From(goqu.T("genres")).
	Where(
		goqu.C("status").Eq(1),
	).
	Select(
		goqu.I("id"),
		goqu.I("name"),
		goqu.I("slug"),
		goqu.I("image"),
		goqu.I("position"),
		goqu.I("status"),
	)
	var genres []models.Genre
	if err := ds.ScanStructs(&genres); err != nil {
		return nil, fmt.Errorf("Faile get all genres:%v", err)
	}

	return genres, nil
}