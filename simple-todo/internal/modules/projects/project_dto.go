package projects

type CreateProjectDTO struct {
	Title string `json:"title"`
	Description string `json:"description"`
	MemberIDs []string `json:"member_ids"`
}

type ProjectResponse struct {
	ID           string   `json:"id"`
	Title        string   `json:"title"`
	Description  string   `json:"description"`
	Status       string   `json:"status"`
	MemberIDs    []string `json:"member_ids"`
}

type UpdateProjectInput struct {
	ID 			string  `json:"id"`
	Title        string   `json:"title"`
	Description  string   `json:"description"`
	Status       string   `json:"status"`
	MemberIDs    []string `json:"member_ids"`
}
