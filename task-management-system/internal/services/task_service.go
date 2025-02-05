package services

import (
	"taskman/internal/models"
	"taskman/internal/repositories"
)

type TaskService interface {
	CreateTask(Task *models.Task) error
	GetAllTasks() ([]models.Task, error)
	GetTaskById(id uint) (*models.Task, error)
	UpdateTask(id uint, task *models.Task) error
	DeleteTask(id uint) error
}

type taskService struct {
	repo repositories.TaskRepository
}

func NewTaskService(repo repositories.TaskRepository) TaskService {
	return &taskService{repo: repo}
}

func (s *taskService) CreateTask(task *models.Task) error {
	return s.repo.Create(task)
}

func (s *taskService) GetAllTasks() ([]models.Task, error) {
	return s.repo.GetAllTasks()
}

func (s *taskService) GetTaskById(id uint) (*models.Task, error) {
	return s.repo.GetTaskById(id)
}

func (s *taskService) UpdateTask(id uint, task *models.Task) error {
	return s.repo.UpdateTask(id, task)
}

func (s *taskService) DeleteTask(id uint) error {
	return s.repo.DeleteTask(id)
}