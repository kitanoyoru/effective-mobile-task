package service

import (
	"context"

	"github.com/kitanoyoru/effective-mobile-task/internal/dtos"
	"github.com/kitanoyoru/effective-mobile-task/internal/models"
	"github.com/kitanoyoru/effective-mobile-task/internal/sessions/cache"
	"github.com/kitanoyoru/effective-mobile-task/internal/sessions/store"
)

type Service struct {
	db    *store.StoreSession
	cache *cache.CacheSession
}

func (s *Service) GetPersonResponse(dto *dtos.PersonGetDTO) (*GetPersonResponse, error) {
	var person *models.Person
	var err error

	if dto.WithFilter {
		person, err = s.getPersonFromDB(dto)
	} else {
		person, err = s.getPersonFromCacheOrDB(dto)
	}

	if err != nil {
		return nil, err
	}

	return NewGetPersonResponseFromModel(person), nil
}

func (s *Service) getPersonFromDB(dto dtos.PersonGetDTO) (*models.Person, error) {
	return s.db.PersonRepository.FindByDTO(dto)
}

func (s *Service) getPersonFromCacheOrDB(dto GetPersonDTO) (*models.Person, error) {
	person, err := s.cache.PersonRepository.GetPersonByID(context.Background(), string(*dto.ID))
	if err != nil {
		person, err = s.db.PersonRepository.FindByID(*dto.ID)
	}
	return person, err
}
