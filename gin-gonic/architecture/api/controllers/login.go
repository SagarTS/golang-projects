package controllers

import (
	"architecture/api/services"
	"architecture/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginUser(c *gin.Context){
	var body models.LoginCredentials

	if c.BindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"error": "failed to read body",
		})
		return
	}

	isUserAlreadyRegistered := services.LoginUser(body.Email, body.Password)

	if !(isUserAlreadyRegistered) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "could not login user",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":"congratulations, you are now logged in",
		})
	}
}