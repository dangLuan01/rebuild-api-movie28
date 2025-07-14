package v1service

import (
	"bytes"
	"encoding/json"

	"github.com/dangLuan01/rebuild-api-movie28/internal/models"
	"github.com/dangLuan01/rebuild-api-movie28/internal/utils"
	"github.com/dangLuan01/rebuild-api-movie28/pkg/search"
	"github.com/elastic/go-elasticsearch/v7"
)

type searchService struct {
	elasticsearch *search.ElasticSearchService
}

func NewSearchService(elasticsearchClient *elasticsearch.Client) SearchService {
	return &searchService{
		elasticsearch: search.NewElasticSearchService(elasticsearchClient),
	}
}

func (ss *searchService) SearchMovie(querySearch string) ([]models.Movie, error) {

	index := utils.GetEnv("ELASTIC_INDEX", "my_elasticsearch")
	var buf bytes.Buffer

	query := map[string]interface{} {
		"query": map[string]interface{}{
			"multi_match": map[string]interface{}{
				"query":  querySearch,
				"fields": []string{"name^2", "origin_name"},
				"fuzziness": "AUTO",
			},
		},
		"size": 50,	
		"from": 0,
	}

	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		return nil, utils.NewError(
			string(utils.ErrCodeInternal),
			"Failed Encode query search.",
		)
	}

	movie, err := ss.elasticsearch.Search(index, buf);
	if err != nil {
		return nil, utils.WrapError(
			string(utils.ErrCodeInternal),
			"Failed search movie.",
			err,
		)
	}

	return movie, nil
}