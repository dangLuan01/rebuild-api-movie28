package app

import (
	v1handler "github.com/dangLuan01/rebuild-api-movie28/internal/handler/v1"
	genrerepository "github.com/dangLuan01/rebuild-api-movie28/internal/repository/genre"
	"github.com/dangLuan01/rebuild-api-movie28/internal/repository/redis"
	"github.com/dangLuan01/rebuild-api-movie28/internal/routes"
	v1routes "github.com/dangLuan01/rebuild-api-movie28/internal/routes/v1"
	v1service "github.com/dangLuan01/rebuild-api-movie28/internal/service/v1"
	"github.com/doug-martin/goqu/v9"
)

type GenreModule struct {
	routes routes.Route
}

func NewGenreModule(DB *goqu.Database, rd redis.RedisRepository) *GenreModule {
	genreRepo := genrerepository.NewSqlGenreRepository(DB)
	genreService := v1service.NewGenreService(genreRepo, rd)
	genreHandler := v1handler.NewGenreHandler(genreService)
	genreRoutes := v1routes.NewGenreRoutes(genreHandler)

	return &GenreModule{
		routes: genreRoutes,
	}
}
func (m *GenreModule) Routes() routes.Route {
	return m.routes
}