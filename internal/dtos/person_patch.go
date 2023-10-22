package dtos

type PersonPatchDTO struct {
	Name       string                   `json:"name"`
	Surname    string                   `json:"surname"`
	Patronymic *string                  `json:"patronymic,omitempty"`
	Age        *int64                   `json:"age,omitempty"`
	Gender     *PersonPatchGenderDTO    `json:"gender,omitempty"`
	Country    *[]PersonPatchCountryDTO `json:"country,omitempty"`
}

type PersonPatchGenderDTO struct {
	Gender      string  `json:"gender"`
	Probability float32 `json:"probability"`
}

type PersonPatchCountryDTO struct {
	CountryID   string  `json:"country_ID"`
	Probability float32 `json:"probability"`
}
