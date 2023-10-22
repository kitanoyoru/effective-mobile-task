package service

import (
	"context"

	"github.com/kitanoyoru/effective-mobile-task/internal/requests"
	"github.com/kitanoyoru/effective-mobile-task/internal/responses"
	"github.com/kitanoyoru/effective-mobile-task/internal/sessions/cache"
	"github.com/kitanoyoru/effective-mobile-task/internal/sessions/store"
)

type Service struct {
	db    *store.StoreSession
	cache *cache.CacheSession
}

func (s *Service) GetPersonResponse(ctx context.Context, request *requests.GetPersonRequest) (*responses.GetPersonResponse, error) {
	person, err := s.db.PersonRepository.Find(ctx, request)
	if err != nil {
		return nil, err
	}

	return responses.NewGetPersonResponseFromModel(person), nil
}

func (s *Service) GetFilterPersonResponse(ctx context.Context, request *requests.GetFilterPersonRequest) (*responses.GetFilterPersonResponse, error) {
	persons, err := s.db.PersonRepository.Filter(ctx, request)
	if err != nil {
		return nil, err
	}

	return responses.NewGetFilterPersonResponseFromModel(persons), nil
}
