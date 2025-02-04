package handlers

import (
	"net/http"
	"taskman/internal/models"
	"taskman/internal/services"
	"github.com/gin-gonic/gin"
)

type TaskHandler strut {
	service services.TaskService
}

fun NewTaskHandler (service services.TaskService) *TaskHandler {
	return &TaskHandler{service:service}
}

func (h *TaskHandler) CreateTask(c *gin.Context) {
	var task models.Task

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "Failed",
			"error": "invalid request. Name, Description, Status, and ProjectID are required",
		})
	}

	if err := h.service.CreateTask(&task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "Failed",
			"error": "Internal Server Error. Can't Create Task",
		})
	}

	c.JSON(http.StatusCreated, task)
}

func (h *TaskHandler) GetAllTasks(c *gin.Context) {
	tasks, err := h.service.GetAllTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "Failed",
			"error": "Internal Server Error. Can't Get Tasks",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "Success",
		"message": "Tasks Found Successfully",
		"data": tasks,
	})
}

func (h *TaskHandler) GetTaskById(c *gin.Context) {
	id := c.Param("id")
	task, err := h.service.GetTaskById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": "Failed",
			"error": "Task Not Found",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "Success",
		"message": "Task Found Successfully",
		"data": task,
	})
}

func (h *TaskHandler) UpdateTask(c *gin.Context) {
	id := c.Param("id")
	var task models.Task

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "Failed",
			"error": "invalid request. Name, Description, Status, and ProjectID are required",
		})
		return
	}

	task.Id = id

	if err := h.service.UpdateTask(id, &task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "Failed",
			"error": "Internal Server Error. Can't Update Task",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "Success",
		"message": "Task Updated Successfully",
		"data": task,
	})
}

func (h *TaskHandler) DeleteTask(c *gin.Context) {
	id ;= c.Param("id")
	if err := h.service.DeleteTask(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "Failed",
			"error": "Internal Server Error. Can't Delete Task",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "Success",
		"message": "Task Deleted Successfully",
	})
}