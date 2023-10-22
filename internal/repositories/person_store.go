package repositories

import (
	"context"
	"fmt"

	"github.com/kitanoyoru/effective-mobile-task/internal/dtos"
	"github.com/kitanoyoru/effective-mobile-task/internal/models"
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

func (r *PersonStoreRepository) Save(person *models.Person) error {
	if err := r.db.Save(person).Error; err != nil {
		return err
	}

	return nil
}

func (r *PersonStoreRepository) FindByID(id int) (*models.Person, error) {
	var person models.Person

	if err := r.db.Where(&models.Person{ID: id}).Take(&person).Error; err != nil {
		return nil, err
	}

	return &person, nil
}

func (r *PersonStoreRepository) FindAll() ([]*models.Person, error) {
	var persons []*models.Person

	if err := r.db.Find(&persons).Error; err != nil {
		return nil, err
	}

	return persons, nil
}

func (r *PersonStoreRepository) DeleteByID(id int) error {
	if err := r.db.Delete(&models.Person{ID: id}).Error; err != nil {
		return err
	}

	r.bus.PublishEvent(context.Background(), PersonDeletedEventTopic, events.PersonDeletedEvent{
		Type: events.PersonDeletedEventType,
		Payload: events.PersonDeletedEventPayload{
			ID: fmt.Sprint(id),
		},
	})

	return nil
}

func (r *PersonStoreRepository) DeleteManyByID(ids []string) error {
	if err := r.db.Where("ID in (?)", ids).Delete(&models.Person{}).Error; err != nil {
		return err
	}

	for _, id := range ids {
		r.bus.PublishEvent(context.Background(), PersonDeletedEventTopic, events.PersonDeletedEvent{
			Type: events.PersonDeletedEventType,
			Payload: events.PersonDeletedEventPayload{
				ID: fmt.Sprint(id),
			},
		})
	}

	return nil
}

func (r *PersonStoreRepository) PatchByID(id int, d *dtos.PersonPatchDTO) error {
	p, err := r.FindByID(id)
	if err != nil {
		return err
	}

	p.MergeWithPatchDTO(d)

	if err := r.Save(p); err != nil {
		return err
	}

	r.bus.PublishEvent(context.Background(), PersonUpdatedEventTopic, events.PersonUpdatedEvent{
		Type: events.PersonUpdatedEventType,
		Payload: events.PersonUpdatedEventPayload{
			ID: fmt.Sprint(id),
		},
	})

	return nil
}
