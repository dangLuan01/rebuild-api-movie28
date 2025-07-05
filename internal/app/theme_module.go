package app

import (
	v1handler "github.com/dangLuan01/rebuild-api-movie28/internal/handler/v1"
	"github.com/dangLuan01/rebuild-api-movie28/internal/repository/redis"
	themerepository "github.com/dangLuan01/rebuild-api-movie28/internal/repository/theme"
	"github.com/dangLuan01/rebuild-api-movie28/internal/routes"
	v1routes "github.com/dangLuan01/rebuild-api-movie28/internal/routes/v1"
	v1service "github.com/dangLuan01/rebuild-api-movie28/internal/service/v1"
	"github.com/doug-martin/goqu/v9"
)

type ThemeModule struct {
	routes routes.Route
}

func NewThemeModule(DB *goqu.Database, rd redis.RedisRepository) *ThemeModule {
	themeRepo := themerepository.NewSqlThemeRepository(DB)
	themeService := v1service.NewThemeService(themeRepo, rd)
	themeHandler := v1handler.NewThemeHandler(themeService)
	themeRoutes := v1routes.NewThemeRoutes(themeHandler)

	return &ThemeModule{
		routes: themeRoutes,
	}
}
func (m *ThemeModule) Routes() routes.Route {
	return m.routes
}