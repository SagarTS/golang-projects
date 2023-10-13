package main

import (
	"gin-crud/initializers"
	"gin-crud/models"
)

func init(){
initializers.LoadEnvVariables()
initializers.ConnectToDB()
}

func main(){
	initializers.DB.AutoMigrate(&models.Post{})
}