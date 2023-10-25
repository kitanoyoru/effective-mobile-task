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

func (r *PersonStoreRepository) Save(ctx context.Context, req *requests.PostPersonRequest) (int, error) {
	person := models.Person{
		Name:       req.Name,
		Surname:    req.Surname,
		Patronymic: null.StringFromPtr(req.Patronymic),
	}

	err := r.db.WithContext(ctx).Save(&person).Error

	r.bus.PublishEvent(ctx, events.PersonPostEventTopic, events.PersonPostEvent{
		Type: events.PersonPostEventType,
		Payload: events.PersonPostEventPayload{
			ID:   person.ID,
			Name: person.Name,
		},
	})

	return person.ID, err
}

func (r *PersonStoreRepository) Find(ctx context.Context, req *requests.GetPersonRequest) (models.Person, error) {
	var person models.Person

	search := models.Person{
		ID: req.ID,
	}

	if err := r.db.WithContext(ctx).Preload("Gender").Preload("Country").Where(&search).Take(&person).Error; err != nil {
		return models.Person{}, err
	}

	r.bus.PublishEvent(ctx, events.PersonGetEventTopic, events.PersonGetEvent{
		Type: events.PersonGetEventType,
		Payload: events.PersonGetEventPayload{
			Person: person,
		},
	})

	return person, nil
}

func (r *PersonStoreRepository) Filter(ctx context.Context, req *requests.GetFilterPersonRequest) ([]models.Person, error) {
	var persons []models.Person

	search := models.Person{
		ID:      req.ID,
		Name:    req.Name,
		Surname: req.Surname,
	}

	offset := (req.Page - 1) * req.Limit

	if err := r.db.WithContext(ctx).Preload("Gender").Preload("Country").Offset(offset).Where(&search).Find(&persons).Limit(req.Limit).Error; err != nil {
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

	r.bus.PublishEvent(ctx, events.PersonDeletedEventTopic, events.PersonDeletedEvent{
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

	if err := r.db.WithContext(ctx).Preload("Gender").Preload("Country").Save(&p).Error; err != nil {
		return err
	}

	r.bus.PublishEvent(ctx, events.PersonUpdatedEventTopic, events.PersonUpdatedEvent{
		Type: events.PersonUpdatedEventType,
		Payload: events.PersonUpdatedEventPayload{
			ID: fmt.Sprint(id),
		},
	})

	return nil
}
