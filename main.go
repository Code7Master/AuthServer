package main

import (
	"DevelopHub/AuthServer/controller"
	"DevelopHub/AuthServer/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDatabase()
}

func main() {
	engine := gin.Default()

	engine.POST("/auth/register", controller.Register)
	engine.POST("/auth/login", controller.Login)

	engine.Run(":9190")
}
