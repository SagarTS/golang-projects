package main

import (
	"architecture/api/routes"
	"architecture/infrastructure"
	"architecture/models"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func init(){
	err := godotenv.Load()
  	if err != nil {
    	log.Fatal("Error loading .env file")
  	}
}

var err error

func main(){
	infrastructure.DB, err = gorm.Open(mysql.Open(infrastructure.DbURL(infrastructure.BuildDBConfig())), &gorm.Config{})
	if err != nil {
		fmt.Println("Status:", err)
	}
	infrastructure.DB.AutoMigrate(&models.UserRegistration{})

	r:= routes.SetupRouter()
	r.Run()
}