package categoryrepository

import (
	"fmt"

	"github.com/dangLuan01/rebuild-api-movie28/internal/models"
	"github.com/doug-martin/goqu/v9"
)

type SqlCategoryRepository struct {
	db *goqu.Database
}

func NewSqlMovRepository(DB *goqu.Database) CategoryRepository {
	return &SqlCategoryRepository{
		db: DB,
	}
}

func (cr *SqlCategoryRepository) FindAll() ([]models.Category, error) {

	var categories []models.Category
	ds := cr.db.From(goqu.T("categories")).
	Where(
		goqu.C("status").Eq(1),
	).
	Select(
		goqu.C("id"),
		goqu.C("name"),
		goqu.C("slug"),
		goqu.C("status"),
	).
	Order(goqu.I("position").Asc()).Limit(10)
	
	if err := ds.ScanStructs(&categories); err != nil {
		return nil, fmt.Errorf("Faile fetch category:%v", err)
	}
	
	return categories, nil
}