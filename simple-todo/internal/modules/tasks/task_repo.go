package tasks

import (
	"context"
	"simple-todo/internal/models"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) CreateTask(ctx context.Context, task models.Task) error {
	return r.db.WithContext(ctx).Create(&task).Error
}

func (r *Repository) UpdateTask(ctx context.Context, task models.Task) error {
	return r.db.WithContext(ctx).Model(&models.Task{}).Where("id = ? AND deletedAt IS NULL", task.ID).Updates(task).Error
}

func (r *Repository) SoftDeleteTask(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Model(&models.Task{}).Where("id = ? AND deleted_at IS NULL", id).Update("deleted_at", gorm.Expr("NOW()")).Error
}

func (r *Repository) GetAllTasks(ctx context.Context) ([]models.Task, error) {
	var tasks []models.Task
	err := r.db.WithContext(ctx).Where("deleted_at IS NULL").Preload("Project").
Find(&tasks).Error
	return tasks, err
}

func (r *Repository) FindById(ctx context.Context, id string) (models.Task, error) {
	var task models.Task
	err := r.db.WithContext(ctx).Where("id = ?", id).Preload("Project").First(&task).Error
	return task, err
}
