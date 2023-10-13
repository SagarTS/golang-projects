package controllers

import (
	"architecture/api/services"
	"architecture/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context){
	var users []models.UserRegistration
	err := services.GetAllUsers(&users)

	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"error": "could not fetch users",
		})
	} else {
		c.JSON(http.StatusOK,gin.H{
			"users": users,
		})
	}
}