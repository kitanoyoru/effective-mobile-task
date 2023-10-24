package repositories

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/kitanoyoru/effective-mobile-task/internal/models"
	"github.com/kitanoyoru/effective-mobile-task/internal/sessions/events"
	"github.com/redis/go-redis/v9"

	log "github.com/sirupsen/logrus"
)

const (
	PersonCachePrefix = "person"

	defaultPersonCacheTTL = 5 * time.Minute
)

type PersonCacheRepository struct {
	client *redis.Client

	bus                    *events.EventBusSession
	personUpdatedCtxCancel context.CancelFunc
	personDeletedCtxcancel context.CancelFunc
	personnGetCtxCancel    context.CancelFunc
}

func NewPersonCacheRepository(client *redis.Client, bus *events.EventBusSession) *PersonCacheRepository {
	personDeletedCtx, personUpdatedCtxCancel := context.WithCancel(context.Background())
	personUpdatedCtx, personDeletedCtxCancel := context.WithCancel(context.Background())
	personGetCtx, personGetCtxCancel := context.WithCancel(context.Background())

	r := &PersonCacheRepository{
		client,
		bus,
		personDeletedCtxCancel,
		personUpdatedCtxCancel,
		personGetCtxCancel,
	}

	go bus.AsyncConsumeEvents(personDeletedCtx, events.PersonDeletedEventTopic, r.onPersonDeletedHandler)
	go bus.AsyncConsumeEvents(personUpdatedCtx, events.PersonUpdatedEventTopic, r.onPersonUpdatedHandler)
	go bus.AsyncConsumeEvents(personGetCtx, events.PersonGetEventTopic, r.onPersonGetHander)

	return r
}

func (r *PersonCacheRepository) GetPersonByID(ctx context.Context, id string) (models.Person, error) {
	cacheKey := r.getCacheKey("id", id)
	personBytes, err := r.client.Get(ctx, cacheKey).Bytes()
	if err != nil {
		return models.Person{}, err
	}

	person := models.Person{}
	if err := json.Unmarshal(personBytes, &person); err != nil {
		return models.Person{}, err
	}

	return person, nil
}

func (r *PersonCacheRepository) SetPersonByID(ctx context.Context, person models.Person) error {
	personBytes, err := json.Marshal(person)
	if err != nil {
		return err
	}

	cacheKey := r.getCacheKey("id", fmt.Sprint(person.ID))
	if err = r.client.Set(ctx, cacheKey, personBytes, defaultPersonCacheTTL).Err(); err != nil {
		return err
	}

	return nil
}

func (r *PersonCacheRepository) DeletePersonByID(ctx context.Context, id string) error {
	cacheKey := r.getCacheKey("id", id)
	if err := r.client.Del(ctx, cacheKey).Err(); err != nil {
		return err
	}

	return nil
}

func (r *PersonCacheRepository) onPersonDeletedHandler(event events.PersonDeletedEvent) {
	if err := r.DeletePersonByID(context.Background(), event.Payload.ID); err != nil {
		log.Debugf("Failed to handle PersonDeletedEvent: %+v", err)
	}
}

func (r *PersonCacheRepository) onPersonUpdatedHandler(event events.PersonUpdatedEvent) {
	if err := r.DeletePersonByID(context.Background(), event.Payload.ID); err != nil {
		log.Debugf("Failed to handle PersonUpdatedEvent: %+v", err)
	}
}

func (r *PersonCacheRepository) onPersonGetHander(event events.PersonGetEvent) {
	if err := r.SetPersonByID(context.Background(), event.Payload.Person); err != nil {
		log.Debugf("Failed to handle PersonGetEvent: %+v", err)
	}
}

func (r *PersonCacheRepository) getCacheKey(category, instance string) string {
	return fmt.Sprintf("%s-%s: %s", PersonCachePrefix, category, instance)
}
