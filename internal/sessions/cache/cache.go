package cache

import (
	"github.com/kitanoyoru/effective-mobile-task/internal/config"
	"github.com/kitanoyoru/effective-mobile-task/internal/repositories"
	"github.com/kitanoyoru/effective-mobile-task/pkg/cache"
	"github.com/redis/go-redis/v9"
)

type CacheSession struct {
	redisClient      *redis.Client
	PersonRepository *repositories.PersonCacheRepository
}

func NewCacheSession(cfg *config.RedisConfig) (*CacheSession, error) {
	client, err := cache.ConnectToRedis(cfg)
	if err != nil {
		return nil, err
	}

	personRepository := repositories.NewPersonCacheRepository(client)

	return &CacheSession{
		redisClient:      client,
		PersonRepository: personRepository,
	}, nil
}

func (c *CacheSession) Close() error {
	return c.redisClient.Close()
}
