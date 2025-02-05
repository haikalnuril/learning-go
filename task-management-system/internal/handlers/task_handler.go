package handlers

import (
	"net/http"
	"strconv"
	"taskman/internal/models"
	"taskman/internal/services"
	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	service services.TaskService
}

func NewTaskHandler(service services.TaskService) *TaskHandler {
	return &TaskHandler{service: service}
}

func (h *TaskHandler) CreateTask(c *gin.Context) {
	var task models.Task

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "Failed",
			"error":  "invalid request. Name, Description, Status, and ProjectID are required",
		})
		return
	}

	if err := h.service.CreateTask(&task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "Failed",
			"error":  "Internal Server Error. Can't Create Task",
		})
		return
	}

	c.JSON(http.StatusCreated, task)
}

func (h *TaskHandler) GetAllTasks(c *gin.Context) {
	tasks, err := h.service.GetAllTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "Failed",
			"error":  "Internal Server Error. Can't Get Tasks",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "Success",
		"message": "Tasks Found Successfully",
		"data":    tasks,
	})
}

func (h *TaskHandler) GetTaskById(c *gin.Context) {
	id := c.Param("id")
	taskID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "Failed",
			"error":  "Invalid Task ID",
		})
		return
	}

	task, err := h.service.GetTaskById(uint(taskID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": "Failed",
			"error":  "Task Not Found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "Success",
		"message": "Task Found Successfully",
		"data":    task,
	})
}

func (h *TaskHandler) UpdateTask(c *gin.Context) {
	id := c.Param("id")
	taskID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "Failed",
			"error":  "Invalid Task ID",
		})
		return
	}

	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "Failed",
			"error":  "invalid request. Name, Description, Status, and ProjectID are required",
		})
		return
	}

	task.ID = uint(taskID)

	if err := h.service.UpdateTask(uint(taskID), &task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "Failed",
			"error":  "Internal Server Error. Can't Update Task",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "Success",
		"message": "Task Updated Successfully",
		"data":    task,
	})
}

func (h *TaskHandler) DeleteTask(c *gin.Context) {
	id := c.Param("id")
	taskID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "Failed",
			"error":  "Invalid Task ID",
		})
		return
	}

	if err := h.service.DeleteTask(uint(taskID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "Failed",
			"error":  "Internal Server Error. Can't Delete Task",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "Success",
		"message": "Task Deleted Successfully",
	})
}