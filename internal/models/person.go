package models

import (
	"github.com/guregu/null"
	"github.com/kitanoyoru/effective-mobile-task/internal/requests"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type PersonGender struct {
	ID string `gorm:"primary_key" json:"id"`

	Gender      string  `gorm:"column:gender;type:TEXT;not null" json:"gender"`
	Probability float32 `gorm:"column:probability;type:FLOAT;not null" json:"probability"`
}

func (p *PersonGender) TableName() string {
	return "Person_Gender"
}

type PersonCountry struct {
	ID string `gorm:"primary_key" json:"id"`

	PersonID int `json:"-"`

	CountryID   string  `gorm:"column:id;type:TEXT;not null" json:"country_id"`
	Probability float32 `gorm:"column:probability;type:FLOAT;not null" json:"probability"`
}

func (p *PersonCountry) TableName() string {
	return "Person_Country"
}

type Person struct {
	ID int `gorm:"primary_key" json:"id"`

	Name    string `gorm:"column:name;type:TEXT;not null" json:"name"`
	Surname string `gorm:"column:surname;type:TEXT;not null" json:"surname"`

	Patronymic null.String `gorm:"column:patronymic;type:TEXT" json:"patronymic,omitempty"`
	Age        null.Int    `gorm:"column:age;type:INT" json:"age,omitempty"`

	GenderID *int         `json:"-"`
	Gender   PersonGender `gorm:"foreignKey:GenderID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"gender"`

	Country []PersonCountry `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"country"`
}

func (p *Person) MergeWithPatchRequest(req *requests.PatchPersonRequest) {
	if req.Name != "" {
		p.Name = req.Name
	}

	if req.Surname != "" {
		p.Surname = req.Surname
	}

	if req.Patronymic != nil {
		p.Patronymic = null.StringFromPtr(req.Patronymic)
	}

	if req.Age != nil {
		p.Age = null.IntFromPtr(req.Age)
	}

	if req.Gender != nil {
		p.Gender.Gender = req.Gender.Gender
		p.Gender.Probability = req.Gender.Probability
	}

	if req.Country != nil {
		var countries []PersonCountry
		for _, countryRequest := range *req.Country {
			country := PersonCountry{
				CountryID:   countryRequest.CountryID,
				Probability: countryRequest.Probability,
			}
			countries = append(countries, country)
		}
		p.Country = countries
	}
}
func (p *Person) TableName() string {
	return "Person"
}
