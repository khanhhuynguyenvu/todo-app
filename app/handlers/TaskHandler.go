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

func (th *TaskHandler) Create(c *gin.Context) {
	task := models.Task{}
	bindErr := c.BindJSON(&task)
	if bindErr != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprint(bindErr))
	}

	task, createErr := th.TaskRepo.Create(task)
	if createErr != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprint(createErr))
	}
	c.JSON(http.StatusCreated, task)
}

func (th *TaskHandler) Select(c *gin.Context) {
	id := c.Param("id")
	if task, err := th.TaskRepo.Select(id); err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprint(err))
	} else {
		c.JSON(http.StatusOK, task)
	}
}

func (th *TaskHandler) All(c *gin.Context) {
	tasks := th.TaskRepo.SelectAll()
	c.JSON(http.StatusOK, tasks)
}

func (th *TaskHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := th.TaskRepo.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprint(err))
	} else {
		c.JSON(http.StatusOK, fmt.Sprintf("Ok"))
	}

}
