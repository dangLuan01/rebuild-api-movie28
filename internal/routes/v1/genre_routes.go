package v1routes

import (
	v1handler "github.com/dangLuan01/rebuild-api-movie28/internal/handler/v1"
	"github.com/gin-gonic/gin"
)

type GenreRoutes struct {
	handler *v1handler.GenreHandler
}

func NewGenreRoutes(handler *v1handler.GenreHandler) *GenreRoutes {
	return &GenreRoutes{
		handler: handler,
	}
}

func (ur *GenreRoutes) Register(r *gin.RouterGroup) {
	genres := r.Group("/genre")
	{
		genres.GET("", ur.handler.GetAllGenres)
		genres.GET("/:slug", ur.handler.GetGenreBySlug)
		// users.POST("", ur.handler.CreateUser)
		// users.PUT("/:uuid", ur.handler.UpdateUser)
		// users.DELETE("/:uuid", ur.handler.DeleteUser)
	}
}