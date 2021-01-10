package repositories

import (
	"github.com/todo-app/app/models"
	"github.com/todo-app/app/utils/db"
	utils "github.com/todo-app/app/utils/erros"
	"github.com/volatiletech/sqlboiler/boil"
	"strconv"
)

type TaskRepoBoiler struct{}

func (trb *TaskRepoBoiler) Create(task *models.Task) (*models.Task, error) {
	err := task.Insert(db.DbConn, boil.Infer())
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (trb *TaskRepoBoiler) UpdateById(id string, taskInp *models.Task) (*models.Task, error) {
	task, findErr := trb.Select(id)
	if findErr != nil {
		utils.PanicError(findErr)
		return nil, findErr
	}
	taskInp.ID = task.ID
	updateErr := taskInp.Update(db.DbConn, boil.Infer())
	if updateErr != nil {
		return nil, updateErr
	}
	return task, nil
}

func (trb *TaskRepoBoiler) SelectAll() ([]*models.Task, error) {
	tasks, err := models.Tasks().All(db.DbConn)
	if err != nil {
		return tasks, err
	}
	return tasks, nil
}

func (trb *TaskRepoBoiler) Select(id string) (*models.Task, error) {
	idConvert, convertErr := strconv.Atoi(id)
	if convertErr != nil {
		return nil, convertErr
	}
	task, err := models.Tasks(models.TaskWhere.ID.EQ(idConvert)).One(db.DbConn)
	if err != nil {
		return task, err
	}
	return task, nil
}

func (trb *TaskRepoBoiler) DeleteById(id string) (*models.Task, error) {
	idConvert, convertErr := strconv.Atoi(id)
	if convertErr != nil {
		return nil, convertErr
	}
	task, err := models.FindTask(db.DbConn, idConvert)
	if err != nil {
		return nil, err
	}
	delErr := task.Delete(db.DbConn)
	if err != nil {
		return nil, delErr
	}
	return task, nil
}

func (trb *TaskRepoBoiler) Delete() error {
	err := models.Tasks().DeleteAll(db.DbConn)
	if err != nil {
		return err
	}
	return nil
}
