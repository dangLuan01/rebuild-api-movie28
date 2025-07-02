package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"strconv"

	"time"

	"github.com/dangLuan01/rebuild-api-movie28/internal/config"
	"github.com/redis/go-redis/v9"
)

type redisRepository struct {
	client *redis.Client
}
var ctx = context.Background()

func NewRedisRepository(cfg config.RedisConfig) RedisRepository {
	indexRedis, _ := strconv.Atoi(cfg.DB)

    client := redis.NewClient(&redis.Options{
        Addr:     fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
        Password: cfg.Password,
        DB:       indexRedis,
    })

    return &redisRepository{
		client: client,
	}
}
func (rd *redisRepository) Set(key string, value any) error {

	data, err := json.Marshal(value)
	if err != nil {
		log.Printf("❌ Error marshaling JSON: %v", err)
		return fmt.Errorf("Error marshaling JSON: %v", err)
	}
	timeExp := time.Duration(rand.Intn(200) + 300) * time.Second
	err = rd.client.Set(ctx, key, data, timeExp).Err()
	if err != nil {
		log.Printf("❌ Error setting cache: %v", err)
	}
	return nil
}
func (rd *redisRepository) Get(key string, dest any) bool {

	val, err := rd.client.Get(ctx, key).Result()
	if err != nil {
		return false
	}

	if err := json.Unmarshal([]byte(val), dest); err != nil {
		log.Printf("❌ Error unmarshaling JSON: %v", err)
		return false
	}

	return true
}
