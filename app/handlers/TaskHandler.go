package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/todo-app/app/interfaces"
	"github.com/todo-app/app/models"
	utils "github.com/todo-app/app/utils/erros"
	"net/http"
)

type TaskHandler struct {
	TaskRepo interfaces.ITaskRepo
}

func (th *TaskHandler) Create(c *gin.Context) {
	taskInp := models.Task{}
	bindErr := c.BindJSON(&taskInp)
	if bindErr != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprint(bindErr))
		return
	}

	task, createErr := th.TaskRepo.Create(&taskInp)
	if createErr != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprint(createErr))
		return
	}
	c.JSON(http.StatusCreated, task)
}

func (th *TaskHandler) Select(c *gin.Context) {
	id := c.Param("id")
	task, err := th.TaskRepo.Select(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprint(err))
		return
	}
	c.JSON(http.StatusOK, task)
}

func (th *TaskHandler) UpdateById(c *gin.Context) {
	id := c.Param("id")
	taskInp := models.Task{}
	bindErr := c.BindJSON(&taskInp)
	if bindErr != nil {
		utils.PanicError(bindErr)
		c.JSON(http.StatusBadRequest, bindErr)
		return
	}
	task, updateErr := th.TaskRepo.UpdateById(id,&taskInp)
	if updateErr != nil {
		c.JSON(http.StatusInternalServerError, updateErr)
		return
	}
	c.JSON(http.StatusOK, task)
}

func (th *TaskHandler) All(c *gin.Context) {
	tasks, err := th.TaskRepo.SelectAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprint(err))
		return
	}
	c.JSON(http.StatusOK, tasks)
}

func (th *TaskHandler) Delete(c *gin.Context) {
	if err := th.TaskRepo.Delete(); err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprint(err))
		return
	}
	c.JSON(http.StatusOK, fmt.Sprintf("Ok"))
}

func (th *TaskHandler) DeleteById(c *gin.Context) {
	id := c.Param("id")
	task, delErr := th.TaskRepo.DeleteById(id)
	if delErr != nil {
		c.JSON(http.StatusInternalServerError, task)
		return
	}
	c.JSON(http.StatusOK, task)
}
