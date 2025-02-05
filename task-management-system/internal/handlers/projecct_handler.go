package handlers

import (
	"net/http"
	"strconv"
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
			"error":  "invalid request. Name, Description, and Status are required",
		})
		return
	}

	if err := h.service.CreateProject(&project); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "Failed",
			"error":  "Internal Server Error. Can't Create Project",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  "Success",
		"message": "Project Created Successfully",
		"data":    project,
	})
}

func (h *ProjectHandler) GetAllProjects(c *gin.Context) {
	projects, err := h.service.GetAllProjects()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "Failed",
			"error":  "Internal Server Error. Can't Get Projects",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "Success",
		"message": "Projects Found Successfully",
		"data":    projects,
	})
}

func (h *ProjectHandler) GetProjectById(c *gin.Context) {
	id := c.Param("id")
	projectID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "Failed",
			"error":  "Invalid Project ID",
		})
		return
	}

	project, err := h.service.GetProjectById(uint(projectID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": "Failed",
			"error":  "Project Not Found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "Success",
		"message": "Project Found Successfully",
		"data":    project,
	})
}

func (h *ProjectHandler) UpdateProject(c *gin.Context) {
	id := c.Param("id")
	projectID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "Failed",
			"error":  "Invalid Project ID",
		})
		return
	}

	var project models.Project
	if err := c.ShouldBindJSON(&project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "Failed",
			"error":  "invalid request. Name, Description, and Status are required",
		})
		return
	}

	project.ID = uint(projectID)

	if err := h.service.UpdateProject(uint(projectID), &project); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "Failed",
			"error":  "Internal Server Error. Can't Update Project",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "Success",
		"message": "Project Updated Successfully",
		"data":    project,
	})
}

func (h *ProjectHandler) DeleteProject(c *gin.Context) {
	id := c.Param("id")
	projectID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "Failed",
			"error":  "Invalid Project ID",
		})
		return
	}

	if err := h.service.DeleteProject(uint(projectID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "Failed",
			"error":  "Internal Server Error. Can't Delete Project",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "Success",
		"message": "Project Deleted Successfully",
	})
}