package requests

type GetFilterPersonRequest struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
