package models

import (
	"github.com/guregu/null"
	"github.com/kitanoyoru/effective-mobile-task/internal/requests"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type PersonGender struct {
	ID string `gorm:"primary_key" json:"id"`

	Gender      string  `gorm:"column:gender;type:TEXT;NOT NULL" json:"gender"`
	Probability float32 `gorm:"column:probability;type:FLOAT;NOT NULL" json:"probability"`
}

func (p *PersonGender) TableName() string {
	return "Person_Gender"
}

type PersonCountry struct {
	ID string `gorm:"primary_key" json:"id"`

	PersonID int `json:"-"`

	CountryID   string  `gorm:"column:id;type:TEXT;NOT NULL" json:"country_id"`
	Probability float32 `gorm:"column:probability;type:FLOAT;NOT NULL" json:"probability"`
}

func (p *PersonCountry) TableName() string {
	return "Person_Country"
}

type Person struct {
	ID int `gorm:"primary_key" json:"id"`

	Name    string `gorm:"column:name;type:TEXT;NOT NULL" json:"name"`
	Surname string `gorm:"column:surname;type:TEXT;NOT NULL" json:"surname"`

	Patronymic null.String `gorm:"column:patronymic;type:TEXT" json:"patronymic,omitempty"`
	Age        null.Int    `gorm:"column:age;type:INT" json:"age,omitempty"`

	GenderID *int          `json:"-"`
	Gender   *PersonGender `gorm:"foreignKey:GenderID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"gender"`

	Country []*PersonCountry `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"country"`
}

func (p *Person) MergeWithPatchRequest(req *requests.PatchPersonRequest) {
	if req.Name != "" {
		p.Name = req.Name
	}

	if req.Surname != "" {
		p.Surname = req.Surname
	}

	if req.Patronymic != "" {
		p.Patronymic = null.StringFrom(req.Patronymic)
	}

	if req.Age != 0 {
		p.Age = null.IntFrom(req.Age)
	}

	if req.Gender != "" && req.GenderProbability != 0.0 {
		if p.Gender == nil {
			p.Gender = &PersonGender{}
			p.Gender.Gender = req.Gender
			p.Gender.Probability = req.GenderProbability
		}
	}

	if req.Country != nil {
		var countries []*PersonCountry
		for _, countryRequest := range req.Country {
			country := &PersonCountry{
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
