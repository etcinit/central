package cache

import (
	"github.com/jacobstr/confer"
	"gopkg.in/redis.v2"
)

// RedisService provides the redis client instance for the application.
type RedisService struct {
	Config *confer.Config `inject:""`
	client *redis.Client
}

// Make creates a new connection to Redis. If it has already been created, it
// returns the existing one.
func (r *RedisService) Make() *redis.Client {
	if r.client != nil {
		return r.client
	}

	r.client = redis.NewTCPClient(&redis.Options{
		Addr:     r.Config.GetString("cache.address"),
		Password: r.Config.GetString("cache.password"),
		DB:       int64(r.Config.GetInt("cache.database")),
	})

	return r.client
}
