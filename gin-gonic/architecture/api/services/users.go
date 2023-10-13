package services

import (
	"architecture/infrastructure"
	"architecture/models"
)

func GetAllUsers(users *[]models.UserRegistration) (error){
	if err := infrastructure.DB.Find(users).Error; err != nil {  // this will automatically populate our users variable in controller
        return err
    }
    return nil
}