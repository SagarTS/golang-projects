package services

import (
	"architecture/infrastructure"
	"architecture/models"
)

func LoginUser(email string, password string) bool{
	var user *models.UserRegistration

	if err := infrastructure.DB.Where("email = ?", email).Where("password = ?", password).First(&user).Error; err != nil {
		return false
	}
	return true
}