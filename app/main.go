package main

import (
	"github.com/gin-gonic/gin"
	"github.com/todo-app/app/route"
	"log"
)

func main() {
	log.Print("Starting Todo app")
	r := gin.Default()
	route.Route(r)
	r.Run()
	//defer db.DbConn.Close()
}
