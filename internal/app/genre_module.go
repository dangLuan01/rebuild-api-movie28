package app

import (
	v1handler "github.com/dangLuan01/rebuild-api-movie28/internal/handler/v1"
	genrerepository "github.com/dangLuan01/rebuild-api-movie28/internal/repository/genre"
	"github.com/dangLuan01/rebuild-api-movie28/internal/routes"
	v1routes "github.com/dangLuan01/rebuild-api-movie28/internal/routes/v1"
	v1service "github.com/dangLuan01/rebuild-api-movie28/internal/service/v1"
)

type GenreModule struct {
	routes routes.Route
	
}

func NewGenreModule(ctx *ModuleContext) *GenreModule {
	genreRepo := genrerepository.NewSqlGenreRepository(ctx.DB)
	genreService := v1service.NewGenreService(genreRepo, ctx.Redis)
	genreHandler := v1handler.NewGenreHandler(genreService)
	genreRoutes := v1routes.NewGenreRoutes(genreHandler)

	return &GenreModule{
		routes: genreRoutes,
	}
}
func (m *GenreModule) Routes() routes.Route {
	return m.routes
}