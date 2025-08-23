package app

import (
	v1handler "github.com/dangLuan01/rebuild-api-movie28/internal/handler/v1"
	movierepository "github.com/dangLuan01/rebuild-api-movie28/internal/repository/movie"
	"github.com/dangLuan01/rebuild-api-movie28/internal/routes"
	v1routes "github.com/dangLuan01/rebuild-api-movie28/internal/routes/v1"
	v1service "github.com/dangLuan01/rebuild-api-movie28/internal/service/v1"
)

type MovieModule struct {
	routes routes.Route
}

func NewMovieModule(ctx *ModuleContext) *MovieModule {
	movieRepo := movierepository.NewSqlMovieRepository(ctx.DB)
	movieService := v1service.NewMovieService(movieRepo, ctx.Redis)
	movieHandler := v1handler.NewMovieHandler(movieService)
	movieRoutes := v1routes.NewMovieRoutes(movieHandler)

	return &MovieModule{
		routes: movieRoutes,
	}
}
func (m *MovieModule) Routes() routes.Route {
	return m.routes
}