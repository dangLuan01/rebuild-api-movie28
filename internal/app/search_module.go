package app

import (
	v1handler "github.com/dangLuan01/rebuild-api-movie28/internal/handler/v1"
	"github.com/dangLuan01/rebuild-api-movie28/internal/routes"
	v1routes "github.com/dangLuan01/rebuild-api-movie28/internal/routes/v1"
	v1service "github.com/dangLuan01/rebuild-api-movie28/internal/service/v1"
	"github.com/elastic/go-elasticsearch/v7"
)

type SearchModule struct {
	routes routes.Route
}

func NewSearchModule(ES *elasticsearch.Client) *SearchModule {
	//search := search.NewElasticSearchService(ES)
	searchService := v1service.NewSearchService(ES)
	searchHandler := v1handler.NewSearchHandler(searchService)
	searchRoutes := v1routes.NewSearchRoutes(searchHandler)

	return &SearchModule{
		routes: searchRoutes,
	}
}
func (m *SearchModule) Routes() routes.Route {
	return m.routes
}