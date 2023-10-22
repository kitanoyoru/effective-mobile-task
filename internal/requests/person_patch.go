package requests

type PatchPersonRequest struct {
	Name       string                       `json:"name"`
	Surname    string                       `json:"surname"`
	Patronymic *string                      `json:"patronymic,omitempty"`
	Age        *int64                       `json:"age,omitempty"`
	Gender     *PatchPersonGenderRequest    `json:"gender,omitempty"`
	Country    *[]PatchPersonCountryRequest `json:"country,omitempty"`
}

type PatchPersonGenderRequest struct {
	Gender      string  `json:"gender"`
	Probability float32 `json:"probability"`
}

type PatchPersonCountryRequest struct {
	CountryID   string  `json:"country_ID"`
	Probability float32 `json:"probability"`
}
