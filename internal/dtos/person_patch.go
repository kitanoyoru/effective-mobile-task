package dtos

type PatchPersonDTO struct {
	Name       string                   `json:"name"`
	Surname    string                   `json:"surname"`
	Patronymic *string                  `json:"patronymic,omitempty"`
	Age        *int64                   `json:"age,omitempty"`
	Gender     *PatchPersonGenderDTO    `json:"gender,omitempty"`
	Country    *[]PatchPersonCountryDTO `json:"country,omitempty"`
}

type PatchPersonGenderDTO struct {
	Gender      string  `json:"gender"`
	Probability float32 `json:"probability"`
}

type PatchPersonCountryDTO struct {
	CountryID   string  `json:"country_ID"`
	Probability float32 `json:"probability"`
}
