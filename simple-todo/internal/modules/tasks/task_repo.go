package tasks

import (
	"context"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) CreateTask(ctx context.Context, task Task) error {
	return r.db.WithContext(ctx).Create(&task).Error
}

func (r *Repository) UpdateTask(ctx context.Context, task Task) error {
	return r.db.WithContext(ctx).Model(&Task{}).Where("id = ? AND deletedAt IS NULL", task.ID).Updates(task).Error
}

func (r *Repository) SoftDeleteTask(ctx context.Context, task Task) error {
	return r.db.WithContext(ctx).Model(&Task{}).Where("id = ? AND deleted_at IS NULL", task.ID).Update("deleted_at", gorm.Expr("NOW()")).Error
}

func (r *Repository) GetAllTasks(ctx context.Context, task Task) ([]Task, error) {
	var tasks []Task
	err := r.db.WithContext(ctx).Where("deleted_at IS NULL").Find(&tasks).Error
	return tasks, err
}

func (r *Repository) FindById(ctx context.Context, id string) (Task, error) {
	var task Task
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&task).Error
	return task, err
}
