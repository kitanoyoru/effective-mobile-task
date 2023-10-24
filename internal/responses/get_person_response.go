package responses

import (
	"github.com/kitanoyoru/effective-mobile-task/internal/models"
)

type GetPersonResponse struct {
	ID int `json:"id"`

	Name    string `json:"name"`
	Surname string `json:"surname"`

	Patronymic string `json:"patronymic,omitempty"`
	Age        int64  `json:"age,omitempty"`

	Gender  *GetPersonGenderResponse    `json:"gender,omitempty"`
	Country []*GetPersonCountryResponse `json:"country,omitempty"`
}

func NewGetPersonResponseFromModel(model models.Person) *GetPersonResponse {
	r := GetPersonResponse{}

	r.ID = model.ID

	r.Name = model.Name
	r.Surname = model.Surname

	if !model.Patronymic.IsZero() {
		r.Patronymic = model.Patronymic.String
	}
	if !model.Age.IsZero() {
		r.Age = model.Age.Int64
	}

	r.Gender = NewGetPersonGenderResponseFromModel(model.Gender)
	if len(model.Country) > 0 {
		for _, c := range model.Country {
			r.Country = append(r.Country, NewGetPersonCountryResponseFromModel(c))

		}
	}

	return &r
}

type GetPersonGenderResponse struct {
	Gender      string  `json:"gender,omitempty"`
	Probability float32 `json:"probability,omitempty"`
}

func NewGetPersonGenderResponseFromModel(model models.PersonGender) *GetPersonGenderResponse {
	r := GetPersonGenderResponse{}

	r.Gender = model.Gender
	r.Probability = model.Probability

	return &r
}

type GetPersonCountryResponse struct {
	CountryID   string  `json:"country_id:"`
	Probability float32 `json:"probability"`
}

func NewGetPersonCountryResponseFromModel(model models.PersonCountry) *GetPersonCountryResponse {
	r := GetPersonCountryResponse{}

	r.CountryID = model.CountryID
	r.Probability = model.Probability

	return &r

}
