package v1routes

import (
	v1handler "github.com/dangLuan01/rebuild-api-movie28/internal/handler/v1"
	"github.com/gin-gonic/gin"
)

type ProxyRoutes struct {
	handler *v1handler.ProxyHandler
}

func NewProxyRoutes(handler *v1handler.ProxyHandler) *ProxyRoutes {
	return &ProxyRoutes{
		handler: handler,
	}
}

func (pr *ProxyRoutes)Register(r *gin.RouterGroup) {
	proxy := r.Group("")
	{
		proxy.GET("", pr.handler.PassHeader)
	}
}