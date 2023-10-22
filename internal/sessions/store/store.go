package store

import (
	"github.com/kitanoyoru/effective-mobile-task/internal/config"
	"github.com/kitanoyoru/effective-mobile-task/internal/repositories"
	"github.com/kitanoyoru/effective-mobile-task/pkg/database"
	"gorm.io/gorm"
)

type StoreSession struct {
	db               *gorm.DB
	PersonRepository *repositories.PersonRepository
}

func NewStoreSession(cfg *config.DatabaseConfig) (*StoreSession, error) {
	db, err := database.ConnectToDB(cfg)
	if err != nil {
		return nil, err
	}

	personRepository := repositories.NewPersonRepository(db)

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
