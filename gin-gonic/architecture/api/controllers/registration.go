package controllers

import (
	"architecture/api/services"
	"architecture/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterUser(c *gin.Context){
	var body models.UserRegistration

	if c.BindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"error": "failed to read body",
		})
		return
	}

	err := services.RegisterUser(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "could not create user"})
	} else {
		c.JSON(http.StatusOK,gin.H{"data": body})
	}

}