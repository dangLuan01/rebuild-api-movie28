package v1routes

import (
	v1handler "github.com/dangLuan01/rebuild-api-movie28/internal/handler/v1"
	"github.com/gin-gonic/gin"
)

type SearchRoutes struct {
	handler *v1handler.SearchHandler
}

func NewSearchRoutes(handler *v1handler.SearchHandler) *SearchRoutes {
	return &SearchRoutes{
		handler: handler,
	}
}

func (sr *SearchRoutes) Register(r *gin.RouterGroup) {
	search := r.Group("/search")
	{
		search.GET("", sr.handler.SeachMovie)
	}
}