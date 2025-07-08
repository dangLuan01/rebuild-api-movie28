package searchrepository

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/dangLuan01/rebuild-api-movie28/internal/models"
	"github.com/elastic/go-elasticsearch/v7"
)

type ESMovieRepository struct {
	es *elasticsearch.Client
}

func NewESMovieRepository(ES *elasticsearch.Client) SearchRepository {
	return &ESMovieRepository {
		es: ES,
	}
}
var r map[string]interface{}

func (sr *ESMovieRepository) Search(index string, search bytes.Buffer) ([]models.Movie, error) {
	// es := elasticsearch.ConnectES()
	// if es == nil {
	// 	log.Printf("Model Elasticsearch: %v", es)
	// 	return entities.SearchResult{}, nil
	// }
	// var buf bytes.Buffer
	// query := map[string]interface{}{
	// 	"query": map[string]interface{}{
	// 		"multi_match": map[string]interface{}{
	// 			"query":  search,
	// 			"fields": []string{"name^2", "origin_name"},
	// 			"fuzziness": "AUTO",
	// 		},
	// 	},
	// 	"size": 50,	
	// 	"from": 0,
	// }
	// if err := json.NewEncoder(&buf).Encode(query); err != nil {
	// 	log.Fatalf("Error encoding query: %s", err)
	// }
	
	res, err := sr.es.Search(
		sr.es.Search.WithContext(context.Background()),
		sr.es.Search.WithIndex(index),
		sr.es.Search.WithBody(&search),
		sr.es.Search.WithTrackTotalHits(true),
		sr.es.Search.WithPretty(),
	)
	if err != nil {

		return nil, fmt.Errorf("Elasticsearch search error:%v", err)
	}
	defer res.Body.Close()
	if res.IsError() {

		return nil, fmt.Errorf("Elasticsearch response error: %s", res.String())	
	}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {

		return nil, fmt.Errorf("Error parsing the response body: %s", err)
	}
	
	hits := r["hits"].(map[string]interface{})["hits"].([]interface{})
	//movie := []entities.Movie{}
	movie := make([]models.Movie, 0)
	for _, hit := range hits {
		hitMap := hit.(map[string]interface{})
		if source, ok := hitMap["_source"].(map[string]interface{}); ok {
			movie = append(movie, models.Movie {
				Name: source["name"].(string),
				Origin_name: source["origin_name"].(string),
				Slug: source["slug"].(string),
				Image: models.Image{
					Poster: source["poster"].(string),
				},	
				Type: source["type"].(string),
				Age: source["age"].(string),
				Release_date: int(source["release_date"].(float64)),
				Runtime: source["runtime"].(string),
			})
		}
	}

	return movie, nil
}