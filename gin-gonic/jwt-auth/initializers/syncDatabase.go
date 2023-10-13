package initializers

import "gin-jwt-auth/models"

func SyncDatabase(){
	DB.AutoMigrate(&models.User{})
}