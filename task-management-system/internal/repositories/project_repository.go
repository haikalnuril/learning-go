package respositories

import (
	"taskman/internal/models"
	"gorm.io/gorm"
)

type projectRepository interface {
	Create(project *models.Project) error
	GetAllProjects() ([]models.Project, error)
	GetProjectById(id uint) (*models.Project, error)
	UpdateProject(id uint, project *models.Project) error
	DeleteProject(id uint) error
}

type projectRepository struct {
	db *gorm.DB
}

func NewProjectRepository(db *gorm.DB) projectRepository {
	return &projectRepository{db: db}
}

func (r *projectRepository) Create(project *models.Project) error {
	return r.db.Create(project).Error
}

func (r *projectRepository) GetAllProjects() ([]models.Project, error) {
	var projects []models.Project
	err := r.db.Find(&projects).Error
	return projects, err
}

func (r *projectRepository) GetProjectById(id uint) (*models.Project, error) {
	var project models.Project
	err := r.db.First(&project, id).Error
	return &project, err
}

func (r *projectRepository) UpdateProject(id uint, project *models.Project) error {
	return  r.db.Model(&models.Project{}).Where("id = ?", id).Updates(project).Error
}

func (r *projectRepository) DeleteProject(id uint) error {
	return r.db.Delete(&models.Project{}, id).Error
}