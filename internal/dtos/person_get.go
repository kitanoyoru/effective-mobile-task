package dtos

type GetPersonDTO struct {
	ID   *int    `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`

	WithFilter bool
}
