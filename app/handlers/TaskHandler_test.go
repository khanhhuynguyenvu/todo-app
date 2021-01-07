package handlers

import (
	"github.com/todo-app/app/models"
	"testing"
)

type xxx struct{}

func (m *xxx) Create(task models.Task) (models.Task, error) {
	return models.Task{}, nil
}

func TestTaskHandler_CreateTask(t *testing.T) {
	//mockX := &xxx{}
	//handler := TaskHandler{
	//	TaskRepo: mockX,
	//}
	//c := &gin.Context{}
	//handler.CreateTask(c)
}
