package app

import (
	v1handler "github.com/dangLuan01/rebuild-api-movie28/internal/handler/v1"
	userrepository "github.com/dangLuan01/rebuild-api-movie28/internal/repository/user"
	"github.com/dangLuan01/rebuild-api-movie28/internal/routes"
	v1routes "github.com/dangLuan01/rebuild-api-movie28/internal/routes/v1"
	v1service "github.com/dangLuan01/rebuild-api-movie28/internal/service/v1"
)

type UserModule struct {
	routes routes.Route
}

func NewUserModule(ctx *ModuleContext) *UserModule {

	userRepo := userrepository.NewSqlUserRepository(ctx.DB)	
	userService := v1service.NewUserService(userRepo)
	UserHandler := v1handler.NewUserHandler(userService)
	userRoutes := v1routes.NewUserRoutes(UserHandler)

	return &UserModule{
		routes: userRoutes,
	}
}
func (m *UserModule) Routes() routes.Route {
	return m.routes
}