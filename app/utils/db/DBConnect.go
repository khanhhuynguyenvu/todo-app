package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	util "github.com/todo-app/app/utils/erros"
)

var DbConn *sql.DB

func init() {
	db, err := sql.Open("mysql", "root2:root2@tcp(127.0.0.1:3306)/tododb")
	util.PanicError(err)
	DbConn = db
}
