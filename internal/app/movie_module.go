package app

import (
	v1handler "github.com/dangLuan01/rebuild-api-movie28/internal/handler/v1"
	movierepository "github.com/dangLuan01/rebuild-api-movie28/internal/repository/movie"
	"github.com/dangLuan01/rebuild-api-movie28/internal/repository/redis"
	"github.com/dangLuan01/rebuild-api-movie28/internal/routes"
	v1routes "github.com/dangLuan01/rebuild-api-movie28/internal/routes/v1"
	v1service "github.com/dangLuan01/rebuild-api-movie28/internal/service/v1"
	"github.com/doug-martin/goqu/v9"
)

type MovieModule struct {
	routes routes.Route
}

func NewMovieModule(DB *goqu.Database, rd redis.RedisRepository) *MovieModule {
	movieRepo := movierepository.NewSqlMovieRepository(DB)
	movieService := v1service.NewMovieService(movieRepo, rd)
	movieHandler := v1handler.NewMovieHandler(movieService)
	movieRoutes := v1routes.NewMovieRoutes(movieHandler)

	return &MovieModule{
		routes: movieRoutes,
	}
}
func (m *MovieModule) Routes() routes.Route {
	return m.routes
}