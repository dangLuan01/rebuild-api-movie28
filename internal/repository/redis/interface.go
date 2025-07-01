package redis

type RedisRepository interface {
	Set(key string, value any)
	Get(key string, dest any) bool
}