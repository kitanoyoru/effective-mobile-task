package requests

type PatchPersonRequest struct {
	Name              string                      `json:"name"`
	Surname           string                      `json:"surname"`
	Patronymic        string                      `json:"patronymic,omitempty"`
	Age               int64                       `json:"age,omitempty"`
	Gender            string                      `json:"gender"`
	GenderProbability float32                     `json:"probability"`
	Country           []PatchPersonCountryRequest `json:"country,omitempty"`
}

type PatchPersonCountryRequest struct {
	CountryID   string  `json:"country_ID"`
	Probability float32 `json:"probability"`
}
