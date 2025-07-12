package v1service

import (
	"bytes"
	"encoding/json"

	"github.com/dangLuan01/rebuild-api-movie28/internal/config"
	"github.com/dangLuan01/rebuild-api-movie28/internal/models"
	searchrepository "github.com/dangLuan01/rebuild-api-movie28/internal/repository/search"
	"github.com/dangLuan01/rebuild-api-movie28/internal/utils"
)

type searchService struct {
	repo searchrepository.SearchRepository
}

func NewSearchService(repo searchrepository.SearchRepository) SearchService {
	return &searchService{
		repo: repo,
	}
}

func (ss *searchService) SearchMovie(querySearch string) ([]models.Movie, error) {

	index := config.NewConfig().ElasticSearch.Index
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

	movie, err := ss.repo.Search(index, buf);
	if err != nil {
		return nil, utils.WrapError(
			string(utils.ErrCodeInternal),
			"Failed search movie.",
			err,
		)
	}

	return movie, nil
}