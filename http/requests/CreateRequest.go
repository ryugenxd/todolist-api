package requests

type CreateRequest struct {
	Title string `json:"title"`
	Description string `json:"description"`
}