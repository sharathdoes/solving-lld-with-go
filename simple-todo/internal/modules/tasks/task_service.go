package tasks

import (
	"context"
	"simple-todo/internal/modules/auth"
	"simple-todo/internal/modules/projects"

	"github.com/google/uuid"
)


type Service struct {
	repo     *Repository
	projRepo *projects.Repository
	userRepo *auth.Repository
}

func NewService(r *Repository, 	projRepo *projects.Repository, 	userRepo *auth.Repository) *Service {
	return &Service{r, projRepo, userRepo}
}

func(s *Service) CreateTask(ctx context.Context, title string, ProjectId string, description string, assignee_ids []string) (*Task, error) {
		project, err := s.projRepo.FindById(ctx, ProjectId)
		if err !=nil {
			return nil, err
		}
		assignees, err:=s.userRepo.FindByIds(ctx, assignee_ids)
		if err !=nil {
			return nil, err
		}
		task:=Task{ ID:uuid.NewString(),Title:title, Description:description, Status:"Pending", ProjectID: ProjectId, Project:project, Assignees: assignees}
		err= s.repo.CreateTask(ctx, task)
		if err!=nil {
			return nil,err
		}
		return &task,nil
}

func (s *Service) GetTasks(ctx context.Context) ( []Task, error) {
		return s.repo.GetAllTasks(ctx)
}