package responses

import "github.com/kitanoyoru/effective-mobile-task/internal/models"

type GetFilterPersonResponse struct {
	Persons []*GetPersonResponse `json:"persons"`
}

func NewGetFilterPersonResponseFromModel(models []models.Person) *GetFilterPersonResponse {
	r := GetFilterPersonResponse{}

	for _, model := range models {
		r.Persons = append(r.Persons, NewGetPersonResponseFromModel(model))
	}

	return &r
}
