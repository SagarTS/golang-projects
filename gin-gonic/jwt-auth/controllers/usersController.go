package controllers

import (
	"gin-jwt-auth/initializers"
	"gin-jwt-auth/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context){
	// Get the email/pass off req Body
	var body struct{
		Email string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	// Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password),10)

	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"error": "Failed to hash password.",
		})
		return
	}

	// Create the user
	user := models.User{Email: body.Email, Password: string(hash)}

	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"error": "Failed to create user.",
		})
	}

	// Respond
	c.JSON(http.StatusOK, gin.H{})
}