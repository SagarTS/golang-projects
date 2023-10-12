package routes

import (
	"architecture/api/controllers"
	"architecture/api/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/register", controllers.RegisterUser)
	r.POST("/login", middlewares.RateLimiter() ,controllers.LoginUser)
	r.GET("/users", controllers.GetAllUsers)

	return r
}