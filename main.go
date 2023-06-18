package main

import (
	"AuthServer/controller"
	"AuthServer/initializers"
	"AuthServer/middleware"

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
	engine.POST("/auth/logout", middleware.RequireAuth, controller.Logout)
	engine.Run(":9190")
}
