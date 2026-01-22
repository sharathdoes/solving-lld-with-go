package projects

import (
	"context"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository (db *gorm.DB) *Repository{
	return &Repository{db}
}

func (r *Repository) CreateProject(ctx context.Context, project *Project) error {
	return r.db.WithContext(ctx).Create(project).Error
}

func (r *Repository) GetProjects(ctx context.Context) ( []Project,error ) {
	var projects []Project
	err:=r.db.WithContext(ctx).Where("deleted_at IS NULL").Preload("Members").Find(&projects).Error
	return projects, err
}


func (r *Repository) UpdateProject(ctx context.Context, project *Project) error {
	return r.db.WithContext(ctx).Model(&Project{}).Where("id = ? AND deleted_at IS NULL", project.ID).Updates(project).Error
}

func (r *Repository) FindById(ctx context.Context, id string) (Project, error) {
    var project Project
	err:=r.db.WithContext(ctx).Where("id = ? AND deleted_at IS NULL", id).First(&project).Error
	return project, err
}


