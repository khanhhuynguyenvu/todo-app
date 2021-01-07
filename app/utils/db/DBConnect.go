package db

import (
	"database/sql"
	util "github.com/todo-app/app/utils/erros"
)

var Conn *sql.DB

func init() {
	db, err := sql.Open("mysql", "root2:root2@tcp(127.0.0.1:3306)/tododb")
	util.PanicError(err)
	Conn = db
}

func DbConn() *sql.DB {
	return Conn
}
