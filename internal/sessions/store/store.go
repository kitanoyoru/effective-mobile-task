package store

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/kitanoyoru/effective-mobile-task/internal/config"
	"github.com/kitanoyoru/effective-mobile-task/internal/models"
	"github.com/kitanoyoru/effective-mobile-task/internal/repositories"
	"github.com/kitanoyoru/effective-mobile-task/internal/sessions/events"
	"github.com/kitanoyoru/effective-mobile-task/pkg/database"
	"gorm.io/gorm"
)

type StoreSession struct {
	db               *gorm.DB
	PersonRepository *repositories.PersonStoreRepository
}

func NewStoreSession(cfg *config.DatabaseConfig, bus *events.EventBusSession) (*StoreSession, error) {
	db, err := database.ConnectToDB(cfg)
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(models.PersonGender{}, models.PersonCountry{})
	db.AutoMigrate(models.Person{})

	personRepository := repositories.NewPersonStoreRepository(db, bus)

	return &StoreSession{
		db,
		personRepository,
	}, nil
}

func (s *StoreSession) Close() error {
	sqlDB, err := s.db.DB()
	if err != nil {
		return err
	}

	return sqlDB.Close()
}
