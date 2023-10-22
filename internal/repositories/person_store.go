package repositories

import (
	"context"
	"fmt"

	"github.com/guregu/null"
	"github.com/kitanoyoru/effective-mobile-task/internal/models"
	"github.com/kitanoyoru/effective-mobile-task/internal/requests"
	"github.com/kitanoyoru/effective-mobile-task/internal/sessions/events"
	"gorm.io/gorm"
)

const (
	PersonDeletedEventTopic = "person:delete"
	PersonUpdatedEventTopic = "person:update"
)

type PersonStoreRepository struct {
	db  *gorm.DB
	bus *events.EventBusSession
}

func NewPersonStoreRepository(db *gorm.DB, bus *events.EventBusSession) *PersonStoreRepository {
	return &PersonStoreRepository{
		db,
		bus,
	}
}

func (r *PersonStoreRepository) Save(ctx context.Context, req *requests.PostPersonRequest) error {
	person := models.Person{
		Name:       req.Name,
		Surname:    req.Surname,
		Patronymic: null.StringFromPtr(req.Patronymic),
	}

	if err := r.db.WithContext(ctx).Save(person).Error; err != nil {
		return err
	}

	return nil
}

func (r *PersonStoreRepository) Find(ctx context.Context, req *requests.GetPersonRequest) (models.Person, error) {
	var person models.Person

	search := models.Person{
		ID: req.ID,
	}

	if err := r.db.WithContext(ctx).Preload("Person_Gender").Preload("Person_Country").Where(&search).Take(&person).Error; err != nil {
		return models.Person{}, err
	}

	return person, nil
}

func (r *PersonStoreRepository) Filter(ctx context.Context, req *requests.GetFilterPersonRequest) ([]models.Person, error) {
	var persons []models.Person

	search := models.Person{
		ID:   req.ID,
		Name: req.Name,
	}

	if err := r.db.WithContext(ctx).Preload("Person_Gender").Preload("Person_Country").Where(&search).Find(&persons).Error; err != nil {
		return nil, err
	}

	return persons, nil
}

func (r *PersonStoreRepository) Delete(ctx context.Context, req *requests.DeletePersonRequest) error {
	search := models.Person{
		ID: req.ID,
	}

	if err := r.db.WithContext(ctx).Delete(&search).Error; err != nil {
		return err
	}

	r.bus.PublishEvent(ctx, PersonDeletedEventTopic, events.PersonDeletedEvent{
		Type: events.PersonDeletedEventType,
		Payload: events.PersonDeletedEventPayload{
			ID: fmt.Sprint(search.ID),
		},
	})

	return nil
}

func (r *PersonStoreRepository) PatchByID(ctx context.Context, id int, req *requests.PatchPersonRequest) error {
	p, err := r.Find(ctx, &requests.GetPersonRequest{ID: id})
	if err != nil {
		return err
	}

	p.MergeWithPatchRequest(req)

	if err := r.db.WithContext(ctx).Save(p).Error; err != nil {
		return err
	}

	r.bus.PublishEvent(ctx, PersonUpdatedEventTopic, events.PersonUpdatedEvent{
		Type: events.PersonUpdatedEventType,
		Payload: events.PersonUpdatedEventPayload{
			ID: fmt.Sprint(id),
		},
	})

	return nil
}
