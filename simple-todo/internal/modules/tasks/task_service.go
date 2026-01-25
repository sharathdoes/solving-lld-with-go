package tasks

import (
	"context"
	"simple-todo/internal/models"
	"simple-todo/internal/modules/auth"

	"github.com/google/uuid"
)


type Service struct {
	repo     *Repository
	userRepo *auth.Repository
}

func NewService(r *Repository, 	userRepo *auth.Repository) *Service {
	return &Service{r, userRepo}
}

func(s *Service) CreateTask(ctx context.Context, title string, ProjectId string, description string, assignee_ids []string) (*models.Task, error) {
		assignees, err:=s.userRepo.FindByIds(ctx, assignee_ids)
		if err !=nil {
			return nil, err
		}
		task:=models.Task{ ID:uuid.NewString(),Title:title, Description:description, Status:"Pending", ProjectID: ProjectId, Assignees: assignees}
		err= s.repo.CreateTask(ctx, task)
		if err!=nil {
			return nil,err
		}
		return &task,nil
}

func (s *Service) GetTasks(ctx context.Context) ( []models.Task, error) {
		return s.repo.GetAllTasks(ctx)
}