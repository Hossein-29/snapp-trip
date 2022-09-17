package main

import (
	"example/snapp/databases"
	"example/snapp/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	databases.ConnectToDatabase()
	router := gin.Default()
	router.GET("/hello/:name", routers.SayHello)
	router.POST("/create/rule", routers.CreateRule)
	router.POST("/create/ticket", routers.CreateTicket)
	router.Run("localhost:8080")
}
