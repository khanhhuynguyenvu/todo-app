package repositories

import (
	"github.com/todo-app/app/models"
	db "github.com/todo-app/app/utils/db"
	utils "github.com/todo-app/app/utils/erros"
)

type ITaskRepo interface {
	Create(task models.Task) (models.Task, error)
	SelectAll() []models.Task
	Select(id string) (models.Task, error)
	Delete(id string) error
}

type TaskRepo struct{}

func (trp *TaskRepo) Create(task models.Task) (models.Task, error) {
	db := db.DbConn
	query := "insert into task(title,content,created_at) values(?,?,?);"
	stmt, stmtErr := db.Prepare(query)
	if stmtErr != nil {
		return task, stmtErr
	}
	res, queryErr := stmt.Exec(task.Title, task.Content, task.CreatedAt)
	if queryErr != nil {
		return task, queryErr
	}

	id, idErr := res.LastInsertId()
	if idErr != nil {
		return task, idErr
	}
	task.Id = id
	return task, nil
}

func (trp *TaskRepo) SelectAll() []models.Task {
	db := db.DbConn
	tasks := make([]models.Task, 0, 10)
	query := "select * from task;"
	rows, queryErr := db.Query(query)
	utils.PanicError(queryErr)
	for rows.Next() {
		var task models.Task
		err := rows.Scan(&task.Id, &task.Title, &task.Content, &task.CreatedAt, &task.UpdatedAt, &task.DeleteAt)
		utils.PanicError(err)
		tasks = append(tasks, task)
	}
	return tasks
}

func (trp *TaskRepo) Select(id string) (models.Task, error) {
	var task models.Task
	db := db.DbConn
	query := "select * from task where id = ?"
	row := db.QueryRow(query, id)
	if err := row.Err(); err != nil {
		return task, err
	}
	scanErr := row.Scan(&task.Id, &task.Title, &task.Content, &task.CreatedAt, &task.UpdatedAt, &task.DeleteAt)
	return task, scanErr
}

func (trp *TaskRepo) Delete(id string) error {
	db := db.DbConn
	query := "delete from task where id = ?;"
	row := db.QueryRow(query, id)
	if err := row.Err(); err != nil {
		return err
	}
	return nil
}
