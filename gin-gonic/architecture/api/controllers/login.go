package controllers

import (
	"architecture/api/services"
	"architecture/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)


func LoginUser(c *gin.Context){
	validate := validator.New()
	var body models.LoginCredentials

	

	if c.BindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"error": "failed to read body",
		})
		return
	}

	err := validate.Struct(body)

	if(err != nil){
		for _,e := range err.(validator.ValidationErrors){
			fmt.Printf("Field: %s, Error: %s\n",e.Field(), e.Tag())	
		}
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