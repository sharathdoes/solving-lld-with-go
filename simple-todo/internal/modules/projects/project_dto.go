package projects

type CreateProject struct {
	Title string `json:"title"`
	Description string `json:"description"`
	OwnerID string `json:"ownerId"`
	MemberIDs []string `json:"member_ids"`
}

type ProjectResponse struct {
	ID           string   `json:"id"`
	Title        string   `json:"title"`
	Description  string   `json:"description"`
	Status       string   `json:"status"`
	MemberIDs    []string `json:"member_ids"`
}
