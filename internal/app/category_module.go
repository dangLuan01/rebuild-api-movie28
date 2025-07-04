package app

import (
	v1handler "github.com/dangLuan01/rebuild-api-movie28/internal/handler/v1"
	categoryrepository "github.com/dangLuan01/rebuild-api-movie28/internal/repository/category"
	"github.com/dangLuan01/rebuild-api-movie28/internal/routes"
	v1routes "github.com/dangLuan01/rebuild-api-movie28/internal/routes/v1"
	v1service "github.com/dangLuan01/rebuild-api-movie28/internal/service/v1"
	"github.com/doug-martin/goqu/v9"
)

type CategoryModule struct {
	routes routes.Route
}

func NewCategoryModule(DB *goqu.Database) *CategoryModule {
	categoryRepo 	:= categoryrepository.NewSqlMovRepository(DB)
	categoryService := v1service.NewCategoryService(categoryRepo)
	categoryHandler := v1handler.NewCategoryHandler(categoryService)
	categoryRoutes 	:= v1routes.NewCategoryRoutes(categoryHandler)

	return &CategoryModule{
		routes: categoryRoutes,
	}
}

func (cr *CategoryModule)Routes() routes.Route  {
	return cr.routes
}