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
	err:=r.db.WithContext(ctx).Find(&projects).Error
	return projects, err
}


func (r *Repository) UpdateProject(ctx context.Context, project *Project) error {
    return r.db.WithContext(ctx).Save(project).Error
}
func (r *Repository) FindById(ctx context.Context, id string) (Project, error) {
    var project Project
	// err:=r.db.WithContext(ctx).Where("ID = ?", id).First(&project).Error
	 err := r.db.WithContext(ctx).First(&project, "id = ?", id).Error

	return project, err
}


