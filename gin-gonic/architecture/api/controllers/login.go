package controllers

import (
	"architecture/api/services"
	"architecture/models"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ApiError struct {
    Field string
    Msg   string
}

func msgForTag(tag string) string {
    switch tag {
    case "required":
        return "This field is required"
    case "email":
        return "Invalid email"
    }
    return ""
}

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
	if err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]ApiError, len(ve))
			for i, fe := range ve {
				out[i] = ApiError{fe.Field(), msgForTag(fe.Tag())}
			}
			c.JSON(http.StatusBadRequest, gin.H{"errors": out})
		}
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