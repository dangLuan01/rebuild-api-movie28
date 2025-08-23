package app

import (
	v1handler "github.com/dangLuan01/rebuild-api-movie28/internal/handler/v1"
	themerepository "github.com/dangLuan01/rebuild-api-movie28/internal/repository/theme"
	"github.com/dangLuan01/rebuild-api-movie28/internal/routes"
	v1routes "github.com/dangLuan01/rebuild-api-movie28/internal/routes/v1"
	v1service "github.com/dangLuan01/rebuild-api-movie28/internal/service/v1"
)

type ThemeModule struct {
	routes routes.Route
}

func NewThemeModule(ctx *ModuleContext) *ThemeModule {
	themeRepo := themerepository.NewSqlThemeRepository(ctx.DB)
	themeService := v1service.NewThemeService(themeRepo, ctx.Redis)
	themeHandler := v1handler.NewThemeHandler(themeService)
	themeRoutes := v1routes.NewThemeRoutes(themeHandler)

	return &ThemeModule{
		routes: themeRoutes,
	}
}
func (m *ThemeModule) Routes() routes.Route {
	return m.routes
}