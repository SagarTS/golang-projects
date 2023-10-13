package main

import (
	"gin-crud/initializers"

	"gin-crud/controllers"

	"github.com/gin-gonic/gin"
)

func init(){
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main(){
	r := gin.Default()
	r.POST("/posts", controllers.PostsCreate)
	r.GET("/posts", controllers.PostsIndex)
	r.GET("/posts/:id", controllers.PostsShow)
	r.PUT("/posts/:id", controllers.PostsUpdate)
	r.DELETE("/posts/:id", controllers.PostsDelete)
	r.Run()
}