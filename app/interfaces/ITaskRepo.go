package interfaces

import "github.com/todo-app/app/models"

type ITaskRepo interface {
	Create(task *models.Task) (*models.Task, error)
	SelectAll() ([]*models.Task, error)
	Select(id string) (*models.Task, error)
	UpdateById(id string,task *models.Task) (*models.Task, error)
	DeleteById(id string) (*models.Task, error)
	Delete() error
}
