package projects

import (
	"context"
	"errors"
	"simple-todo/internal/modules/auth"

	"github.com/google/uuid"
)

type Service struct {
	repo     *Repository
	userRepo auth.Repository
}

func NewService(r *Repository, userRepo auth.Repository) *Service {
	return &Service{r, userRepo}
}

func (s *Service) CreateProject(ctx context.Context, title string, description string, OwnerID string, member_ids []string) (*Project, error) {
	owner, err := s.userRepo.FindByID(ctx, OwnerID)
	if err != nil {
		return nil, errors.New("OWNER_NOT_FOUND")
	}
	members, err := s.userRepo.FindByIds(ctx, member_ids)
	if err != nil {
		return nil, errors.New("MEMBER_NOT_FOUND")
	}
	proj := &Project{ID: uuid.NewString(), Title: title, Description: description, Status: "Pending", OwnerID: owner.ID, Owner: *owner, Members: members}
	if err := s.repo.CreateProject(ctx, proj); err != nil {
		return nil, err
	}
	return proj, nil
}

func (s *Service) UpdateProject(
    ctx context.Context,
    projectID string,
    title string,
    description string,
    ownerID string,
    memberIDs []string,
) (*Project, error) {

    project, err := s.repo.FindByIdWithMembers(ctx, projectID)
    if err != nil {
        return nil, err
    }

    if title != "" {
        project.Title = title
    }

    if description != "" {
        project.Description = description
    }

    if err := s.repo.UpdateProject(ctx, project); err != nil {
        return nil, err
    }

    if len(memberIDs) > 0 {
        members, err := s.userRepo.FindByIds(ctx, memberIDs)
        if err != nil {
            return nil, err
        }

        if err := s.repo.AppendMembers(ctx, project.ID, members); err != nil {
            return nil, err
        }
    }

    return project, nil
}


func (s *Service) GetProjects(ctx context.Context) ([]Project, error) {
	return s.repo.GetProjects(ctx)
}

func (s *Service) GetProjectsWithTasks(ctx context.Context) ([]Project, error) {
	return s.repo.GetProjectsWithTasks(ctx)
}


// func (s *Service) FindMyProjects(ctx context.Context,) ([]Project,error){

// }