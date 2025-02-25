package repositories

import (
	"taskman/internal/models"
	"gorm.io/gorm"
)

type TaskRepository interface {
	Create(task *models.Task) error
	GetAllTasks() ([]models.Task, error)
	GetTaskById(id uint) (*models.Task, error)
	UpdateTask(id uint, task *models.Task) error
	DeleteTask(id uint) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) Create(task *models.Task) error {
	return r.db.Create(task).Error
}

func (r *taskRepository) GetAllTasks() ([]models.Task, error) {
	var tasks []models.Task
	err := r.db.Find(&tasks).Error
	return tasks, err
}

func (r *taskRepository) GetTaskById(id uint) (*models.Task, error) {
	var task models.Task
	err := r.db.First(&task, id).Error
	return &task, err
}

func (r *taskRepository) UpdateTask(id uint, task *models.Task) error {
	return r.db.Model(&models.Task{}).Where("id = ?", id).Updates(task).Error
}

func (r *taskRepository) DeleteTask(id uint) error {
	return r.db.Delete(&models.Task{}, id).Error
}