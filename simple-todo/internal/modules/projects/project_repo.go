package projects

import (
	"context"
	"simple-todo/internal/models"
	"gorm.io/gorm"
)

type Project = models.Project 

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) CreateProject(ctx context.Context, project *models.Project) error {
	return r.db.WithContext(ctx).Create(project).Error
}

func (r *Repository) GetProjects(ctx context.Context) ([]models.Project, error) {
	var projects []models.Project
	err := r.db.WithContext(ctx).Where("deleted_at IS NULL").Preload("Members").Find(&projects).Error
	return projects, err
}

func (r *Repository) GetProjectsWithTasks(ctx context.Context) ([]models.Project, error) {
	var projects []models.Project
	err := r.db.WithContext(ctx).Where("deleted_at IS NULL").Preload("Tasks").Preload("Members").Find(&projects).Error
	return projects, err
}

func (r *Repository) UpdateProject(ctx context.Context, project *models.Project) error {
	return r.db.WithContext(ctx).Model(&Project{}).Where("id = ? AND deleted_at IS NULL", project.ID).Updates(project).Error
}

func (r *Repository) FindById(ctx context.Context, id string) (*models.Project, error) {
	var project models.Project
	err := r.db.WithContext(ctx).Where("id = ? AND deleted_at IS NULL", id).First(&project).Error
	return &project, err
}

func (r *Repository) FindByIdWithMembersAndTasks(
	ctx context.Context,
	id string,
) (*models.Project, error) {
	var project Project
	err := r.db.
		WithContext(ctx).
		Preload("Owner").
		Preload("Members").
		Preload("Tasks").
		First(&project, "id = ? AND deleted_at IS NULL", id).
		Error
	return &project, err
}

func (r *Repository) FindByIdWithMembers(
	ctx context.Context,
	id string,
) (*models.Project, error) {
	var project Project
	err := r.db.
		WithContext(ctx).
		Preload("Owner").
		Preload("Members").
		First(&project, "id = ? AND deleted_at IS NULL", id).
		Error
	return &project, err
}

func (r *Repository) FindMyProjects(ctx context.Context, id string) ([]models.Project, error) {
	var projects []models.Project
	err := r.db.WithContext(ctx).Where("owner_id = ?", id).Find(&projects).Error
	if err != nil {
		return nil, err
	}
	return projects, nil
}

func (r *Repository) AppendMembers(
	ctx context.Context,
	projectID string,
	members []models.User,
) error {
	return r.db.
		WithContext(ctx).
		Model(&models.Project{ID: projectID}).
		Association("Members").
		Append(members)
}
