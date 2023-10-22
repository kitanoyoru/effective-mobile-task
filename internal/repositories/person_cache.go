package repositories

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/kitanoyoru/effective-mobile-task/internal/models"
	"github.com/redis/go-redis/v9"
)

const (
	PersonCachePrefix = "person"

	defaultPersonCacheTTL = 5 * time.Minute
)

type PersonCacheRepository struct {
	client *redis.Client
}

func NewPersonCacheRepository(client *redis.Client) *PersonCacheRepository {
	return &PersonCacheRepository{
		client,
	}
}

func (r *PersonCacheRepository) GetPersonByID(ctx context.Context, id string) (*models.Person, error) {
	cacheKey := r.getCacheKey("id", id)
	personBytes, err := r.client.Get(ctx, cacheKey).Bytes()
	if err != nil {
		return nil, err
	}

	person := models.Person{}
	if err := json.Unmarshal(personBytes, &person); err != nil {
		return nil, err
	}

	return &person, nil
}

func (r *PersonCacheRepository) SetPersonByID(ctx context.Context, id string, person *models.Person) error {
	personBytes, err := json.Marshal(person)
	if err != nil {
		return err
	}

	cacheKey := r.getCacheKey("id", id)
	if err = r.client.Set(ctx, cacheKey, personBytes, defaultPersonCacheTTL).Err(); err != nil {
		return err
	}

	return nil
}

func (r *PersonCacheRepository) getCacheKey(category, instance string) string {
	return fmt.Sprintf("%s-%s: %s", PersonCachePrefix, category, instance)
}
