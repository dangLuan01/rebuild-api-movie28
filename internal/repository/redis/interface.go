package redis

type RedisRepository interface {
	Set(key string, value any) error
	Get(key string, dest any) bool
}