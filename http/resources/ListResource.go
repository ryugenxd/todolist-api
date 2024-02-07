package resources

type ListResource struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}
