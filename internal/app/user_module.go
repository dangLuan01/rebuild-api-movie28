package app

import (
	v1handler "github.com/dangLuan01/rebuild-api-movie28/internal/handler/v1"
	"github.com/dangLuan01/rebuild-api-movie28/internal/repository/redis"
	userrepository "github.com/dangLuan01/rebuild-api-movie28/internal/repository/user"
	"github.com/dangLuan01/rebuild-api-movie28/internal/routes"
	v1routes "github.com/dangLuan01/rebuild-api-movie28/internal/routes/v1"
	v1service "github.com/dangLuan01/rebuild-api-movie28/internal/service/v1"
	"github.com/doug-martin/goqu/v9"
)

type UserModule struct {
	routes routes.Route
}

func NewUserModule(DB *goqu.Database, rd redis.RedisRepository) *UserModule {

	userRepo := userrepository.NewSqlUserRepository(DB)
	userService := v1service.NewUserService(userRepo, rd)
	UserHandler := v1handler.NewUserHandler(userService)
	userRoutes := v1routes.NewUserRoutes(UserHandler)

	return &UserModule{
		routes: userRoutes,
	}
}
func (m *UserModule) Routes() routes.Route {
	return m.routes
}