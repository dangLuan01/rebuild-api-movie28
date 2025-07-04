package v1routes

import (
	v1handler "github.com/dangLuan01/rebuild-api-movie28/internal/handler/v1"
	"github.com/gin-gonic/gin"
)

type CategoryRoutes struct {
	handler *v1handler.CategoryHandler
}

func NewCategoryRoutes(hanler *v1handler.CategoryHandler) *CategoryRoutes {
	return &CategoryRoutes{
		handler: hanler,
	}
}

func (cr *CategoryRoutes) Register(r *gin.RouterGroup) {
	categories := r.Group("/category")
	{
		categories.GET("", cr.handler.GetAllCategory)
	}
}