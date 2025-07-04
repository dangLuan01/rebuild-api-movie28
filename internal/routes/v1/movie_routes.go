package v1routes

import (
	v1handler "github.com/dangLuan01/rebuild-api-movie28/internal/handler/v1"
	"github.com/gin-gonic/gin"
)

type MovieRoutes struct {
	handler *v1handler.MovieHandler
}

func NewMovieRoutes(handler *v1handler.MovieHandler) *MovieRoutes {
	return &MovieRoutes{
		handler: handler,
	}
}

func (mr *MovieRoutes) Register(r *gin.RouterGroup) {
	movies := r.Group("/movie")
	{
		movies.GET("/hot", mr.handler.GetMovieWithHot)
		movies.GET("", mr.handler.GetAllMovies)
		movies.GET("/:slug", mr.handler.GetMovieDetail)
	}
}