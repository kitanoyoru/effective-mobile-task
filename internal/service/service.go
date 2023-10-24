package service

import (
	"context"
	"fmt"

	"github.com/kitanoyoru/effective-mobile-task/internal/models"
	"github.com/kitanoyoru/effective-mobile-task/internal/requests"
	"github.com/kitanoyoru/effective-mobile-task/internal/responses"
	"github.com/kitanoyoru/effective-mobile-task/internal/sessions/cache"
	"github.com/kitanoyoru/effective-mobile-task/internal/sessions/store"
)

type Service struct {
	db    *store.StoreSession
	cache *cache.CacheSession
}

func NewService(db *store.StoreSession, cache *cache.CacheSession) *Service {
	return &Service{
		db,
		cache,
	}
}

func (s *Service) GetPerson(ctx context.Context, request *requests.GetPersonRequest) (*responses.GetPersonResponse, error) {
	var person models.Person
	var err error

	person, err = s.cache.PersonRepository.GetPersonByID(ctx, fmt.Sprint(request.ID))
	if err != nil {
		person, err = s.db.PersonRepository.Find(ctx, request)
		if err != nil {
			return nil, err
		}
	}

	return responses.NewGetPersonResponseFromModel(person), nil
}

func (s *Service) FilterPerson(ctx context.Context, request *requests.GetFilterPersonRequest) (*responses.GetFilterPersonResponse, error) {
	persons, err := s.db.PersonRepository.Filter(ctx, request)
	if err != nil {
		return nil, err
	}

	return responses.NewGetFilterPersonResponseFromModel(persons), nil
}

func (s *Service) AddPerson(ctx context.Context, request *requests.PostPersonRequest) (*responses.PostPersonResponse, error) {
	id, err := s.db.PersonRepository.Save(ctx, request)
	if err != nil {
		return nil, err
	}

	return &responses.PostPersonResponse{
		ID: id,
	}, nil
}

func (s *Service) DeletePerson(ctx context.Context, request *requests.DeletePersonRequest) (*responses.DeletePersonResponse, error) {
	if err := s.db.PersonRepository.Delete(ctx, request); err != nil {
		return nil, err
	}

	return &responses.DeletePersonResponse{
		ID: request.ID,
	}, nil
}

func (s *Service) PatchPerson(ctx context.Context, request *requests.PatchPersonRequest) (*responses.PatchPersonResponse, error) {
	if err := s.db.PersonRepository.PatchByID(ctx, request.ID, request); err != nil {
		return nil, err
	}

	return &responses.PatchPersonResponse{
		ID: request.ID,
	}, nil
}
