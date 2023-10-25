package requests

type DeletePersonRequest struct {
	ID int `json:"id"`
}

type GetFilterPersonRequest struct {
	ID      int    `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Surname string `json:"surname,omitempty"`

	Limit int `json:"limit,omitempty"`
	Page  int `json:"page,omitempty"`
}

type GetPersonRequest struct {
	ID int `json:"id,omitempty"`
}

type PatchPersonRequest struct {
	ID                int                         `json:"id,omitempty"`
	Name              string                      `json:"name"`
	Surname           string                      `json:"surname"`
	Patronymic        string                      `json:"patronymic,omitempty"`
	Age               int64                       `json:"age,omitempty"`
	Gender            string                      `json:"gender"`
	GenderProbability float32                     `json:"probability"`
	Country           []PatchPersonCountryRequest `json:"country,omitempty"`
}

type PatchPersonCountryRequest struct {
	CountryID   string  `json:"country_id"`
	Probability float32 `json:"probability"`
}

type PostPersonRequest struct {
	Name       string  `json:"name"`
	Surname    string  `json:"surname"`
	Patronymic *string `json:"patronymic,omitempty"`
}
