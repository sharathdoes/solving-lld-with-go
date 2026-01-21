package tasks

type CreateTask struct {
	Title string `json:"title"`
	Description string `json:"description"`
	ProjectId string `json:"project_id"`
	AssigneeIDs []string `json:"assignee_ids"`
}


type TaskResponse struct {
	ID     string   `json:"id"`
	Title  string   `json:"title"`
	Status string   `json:"status"`
}