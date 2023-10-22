package models

import (
	"github.com/guregu/null"
	"github.com/kitanoyoru/effective-mobile-task/internal/dtos"
)

type Person struct {
	ID int `gorm:"primary_key;column:id;type:INT;AUTO_INCREMENT" json:"id"`

	Name    string `gorm:"column:name;type:TEXT;not null;" json:"name"`
	Surname string `gorm:"column:surname;type:TEXT;not null;" json:"surname"`

	Patronymic null.String `gorm:"column:patronymic;type:TEXT;" json:"patronymic,omitempty"`
	Age        null.Int    `gorm:"column:age;type:INT;" json:"age,omitempty"`

	Gender  *PersonGender    `gorm:"foreignKey:PersonID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"gender;"`
	Country []*PersonCountry `gorm:"foreignKey:PersonID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"country;"`
}

func (p *Person) MergeWithPatchDTO(d *dtos.PatchPersonDTO) {
	if d.Name != "" {
		p.Name = d.Name
	}

	if d.Surname != "" {
		p.Surname = d.Surname
	}

	if d.Patronymic != nil {
		p.Patronymic = null.StringFromPtr(d.Patronymic)
	}

	if d.Age != nil {
		p.Age = null.IntFromPtr(d.Age)
	}

	if d.Gender != nil {
		if p.Gender == nil {
			p.Gender = &PersonGender{}
		}
		p.Gender.Gender = d.Gender.Gender
		p.Gender.Probability = d.Gender.Probability
	}

	if d.Country != nil {
		var countries []*PersonCountry
		for _, countryDTO := range *d.Country {
			country := &PersonCountry{
				CountryID:   countryDTO.CountryID,
				Probability: countryDTO.Probability,
			}
			countries = append(countries, country)
		}
		p.Country = countries
	}
}

func (p *Person) TableName() string {
	return "Person"
}

type PersonGender struct {
	ID string `gorm:"primary_key;column:id;type:INT;AUTO_INCREMENT;" json:"id"`

	Gender      string  `gorm:"column:gender;type:TEXT;not null;" json:"gender"`
	Probability float32 `gorm:"column:probability;type:FLOAT;not null;" json:"probability"`

	PersonID string `gorm:"column:personID;type:TEXT;" json:"personID"`
}

func (p *PersonGender) TableName() string {
	return "Person_Gender"
}

type PersonCountry struct {
	ID string `gorm:"primary_key;column:id;type:INT;AUTO_INCREMENT;" json:"id"`

	CountryID   string  `gorm:"column:id;type:TEXT;not null;" json:"country_id"`
	Probability float32 `gorm:"column:probability;type:FLOAT;not null;" json:"probability"`

	PersonID string `gorm:"column:personID;type:TEXT;" json:"personID"`
}

func (p *PersonCountry) TableName() string {
	return "Person_Country"
}
