package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/todo-app/app/models"
	repo "github.com/todo-app/app/repositories"
	"net/http"
)

type TaskHandler struct {
	TaskRepo repo.ITaskRepo
}

func (h *TaskHandler) CreateTask(c *gin.Context) {
	task := models.Task{}
	bindErr := c.BindJSON(&task)
	if bindErr != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprint(bindErr))
	}

	task, createErr := h.TaskRepo.Create(task)
	if createErr != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprint(createErr))
	}
	c.JSON(http.StatusCreated, task)
}

func SelectTask(c *gin.Context) {
	id := c.Param("id")
	if task, err := repo.Select(id); err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprint(err))
	} else {
		c.JSON(http.StatusOK, task)
	}
}

func AllTask(c *gin.Context) {
	tasks := repo.SelectAll()
	c.JSON(http.StatusOK, tasks)
}

func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	if err := repo.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprint(err))
	} else {
		c.JSON(http.StatusOK, fmt.Sprintf("Ok"))
	}

}

type TaskSQ struct{}

func (tsq *TaskSQ) handling(c *gin.Context) {
	c.JSON(http.StatusOK, "Ok")
}
