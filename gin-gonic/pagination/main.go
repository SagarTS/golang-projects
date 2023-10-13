package main

import (
	"go-pagination/controllers"
	"go-pagination/initializers"

	"github.com/gin-gonic/gin"
)

func init(){
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDB()
	initializers.CreatePeople()
}

func main(){
	app := gin.Default()

	// Configure app
	app.LoadHTMLGlob("templates/**/*")
	app.Static("/assets", "./assets")

	// Routing
	app.GET("/people", controllers.PeopleIndexGET)
	app.GET("/people/page/:page", controllers.PeopleIndexGET)

	// Start app
	app.Run()
}