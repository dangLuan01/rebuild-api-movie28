package app

import (
	v1handler "github.com/dangLuan01/rebuild-api-movie28/internal/handler/v1"
	"github.com/dangLuan01/rebuild-api-movie28/internal/routes"
	v1routes "github.com/dangLuan01/rebuild-api-movie28/internal/routes/v1"
	v1service "github.com/dangLuan01/rebuild-api-movie28/internal/service/v1"
)
type ProxyModule struct {
	routes routes.Route
}

func NewProxyModule() *ProxyModule {
	
	proxyService := v1service.NewProxyService()
	proxyHandler := v1handler.NewProxyHandler(proxyService)
	proxyRoutes := v1routes.NewProxyRoutes(proxyHandler)

	return &ProxyModule{
		routes: proxyRoutes,
	}
}
func (p *ProxyModule) Routes() routes.Route {
	return p.routes
}