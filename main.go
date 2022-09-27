package main

import (
	"example/snapp/databases"
	"example/snapp/routers"
	"example/snapp/validations"

	"github.com/gin-gonic/gin"
)

func main() {
	databases.ConnectToRedis()
	databases.ConnectToPostgres()
	validations.PreValidation()
	router := gin.Default()
	router.GET("/hello/:name", routers.SayHello)
	router.POST("/create/rule", routers.CreateRule)
	router.POST("/create/ticket", routers.CreateTicket)
	router.Run("localhost:8080")
}
