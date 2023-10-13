package services

import (
	"architecture/infrastructure"
	"architecture/models"
)

func RegisterUser(body *models.UserRegistration)(err error){
	if err = infrastructure.DB.Create(&body).Error; err != nil {
		return err
	}
	return nil
}