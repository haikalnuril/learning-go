package handlers

import (
	"net/http"
	"taskman/internal/models"
	"taskman/internal/services"
	"github.com/gin-gonic/gin"
)

type ProjectHandler struct {
	service services.ProjectService
}

func NewProjectHandler(service services.ProjectService) *ProjectHandler {
	return &ProjectHandler{service: service}
}

func (h *ProjectHandler) CreateProject(c *gin.Context) {
	var project models.Project

	if err := c.ShouldBindJSON(&project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "Failed",
			"error": "invalid request. Name, Description, and Status are required",
		})
	}

	if err := h.service.CreateProject(&project); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "Failed",
			"error": "Internal Server Error. Can't Create Project",
		})
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": "Success",
		"message": "Project Created Successfully",
		"data": project,
	})
}

func (h *ProjectHandler) GetAllProjects(c *gin.Context) {
	projects, err := h.service.GetAllProjects()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "Failed",
			"error": "Internal Server Error. Can't Get Projects",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "Success",
		"message": "Projects Found Successfully",
		"data": projects,
	})
}

func (h *ProjectHandler) GetProjectById(c *gin.Context) {
	id := c.Param("id")
	project, err := h.service.GetProjectById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": "Failed",
			"error": "Project Not Found",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "Success",
		"message": "Project Found Successfully",
		"data": project,
	})
}

func (h *ProjectHandler) UpdateProject(c *gin.Context) {
	id := c.Param("id")
	var project models.Project

	if err := c.ShouldBindJSON(&project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "Failed",
			"error": "invalid request. Name, Description, and Status are required",
		})
	}

	project.ID = id

	if err := h.service.UpdateProject(id, &project); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "Failed",
			"error": "Internal Server Error. Can't Update Project",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "Success",
		"message": "Project Updated Successfully",
		"data": project,
	})
}

func (h *ProjectHandler) DeleteProject(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.DeleteProject(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "Failed",
			"error": "Internal Server Error. Can't Delete Project",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "Success",
		"message": "Project Deleted Successfully",
	})
}