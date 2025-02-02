package services

import (
	"taskman/internal/models"
	"taskman/internal/repositories"
)

type ProjectService interface {
	CreateProject(project *models.Project) error
	GetAllProjects() ([]models.Project, error)
	GetProjectById(id uint) (*models.Project, error)
	UpdateProject(project *models.Project) error
	DeleteProject(id uint) error
}

type projectService struct {
	repo repositories.ProjectRepository
}

func NewProjectService(repo repositories.ProjectRepository) ProjectService {
	return &projectService{repo: repo}
}

func (s *projectService) CreateProject(project *models.Project) error {
	return s.repo.Create(project)
}

func (s *projectService) GetAllProjects() ([]models.Project, error) {
	return s.repo.GetAllProjects()
}

func (s *projectService) GetProjectById(id uint) (*models.Project, error) {
	return s.repo.GetProjectById(id)
}

func (s *projectService) UpdateProject(project *models.Project) error {
	return s.repo.UpdateProject(project)
}

func (s *projectService) DeleteProject(id uint) error {
	return s.repo.DeleteProject(id)
}