package v1routes

import (
	v1handler "github.com/dangLuan01/rebuild-api-movie28/internal/handler/v1"
	"github.com/gin-gonic/gin"
)

type ThemeRoutes struct {
	handler *v1handler.ThemeHandler
}

func NewThemeRoutes(handler *v1handler.ThemeHandler) *ThemeRoutes {
	return &ThemeRoutes{
		handler: handler,
	}
}

func (tr *ThemeRoutes) Register(r *gin.RouterGroup) {
	theme := r.Group("/theme")
	{
		theme.GET("", tr.handler.GetAllThemes)
	}
}