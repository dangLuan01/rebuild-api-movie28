package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/dangLuan01/rebuild-api-movie28/internal/utils"
	"github.com/elastic/go-elasticsearch/v7"
)

type ElasticSearchConfig struct {
	Host  string
	Port  int
	Index string
}

func NewElasticSearchClient() *elasticsearch.Client {
	cfg := ElasticSearchConfig {
		Host:  utils.GetEnv("ELASTIC_HOST", "http://localhost"),
		Port:  utils.GetIntEnv("ELASTIC_PORT", 9200),
		//Index: utils.GetEnv("ELASTIC_INDEX", "my_elasticsearch"),
	}

	client, err := elasticsearch.NewClient(
		elasticsearch.Config{
			Addresses: []string {
				fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
			},
			CompressRequestBody: true,
		},
	)

	if err != nil {
		log.Fatalf("❌ Faile to connecting ES:%s", err)
    }
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
	if _, err = client.Info(
		client.Info.WithContext(ctx),
	); err != nil {
		log.Fatal("❌ Faile ping ES...",)
	}

	log.Println("✅ Connected ES.")

	return client
}
