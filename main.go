package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yourname/go-task-tracker/db"
	"github.com/yourname/go-task-tracker/routes"
)

func main() {
	db.Connect()

	r := gin.Default()
	routes.RegisterRoutes(r)

	r.Run(":8081") // server on localhost:8080
}
