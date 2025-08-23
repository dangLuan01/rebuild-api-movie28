package routes

import (
	"github.com/dangLuan01/rebuild-api-movie28/internal/middleware"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

type Route interface {
	Register(r *gin.RouterGroup)
}

func RegisterRoute(r *gin.Engine, routes ...Route) {
	api 	:= r.Group("/api/v1")
	

	api.Use(
		middleware.ApiKeyMiddleware(),
		middleware.RateLimiterMiddleware(),
		//middleware.AuthMiddleware(),
	)
	api.Use(
		gzip.Gzip(gzip.DefaultCompression),
	)

	for _, route := range routes {
		route.Register(api)
		
	}
}
func RegisterPublicRoute(r *gin.Engine, routes ...Route) {
	public := r.Group("/proxy")

	public.Use(
		middleware.RateLimiterMiddleware(),
	)

	for _, route := range routes {
		route.Register(public)
	}
}