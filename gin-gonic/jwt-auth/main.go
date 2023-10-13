package main

import (
	"gin-jwt-auth/controllers"
	"gin-jwt-auth/initializers"
	middleware "gin-jwt-auth/milddleware"

	"github.com/gin-gonic/gin"
)

func init(){
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main(){
	r := gin.Default()
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth ,controllers.Validate)

	r.Run()
}