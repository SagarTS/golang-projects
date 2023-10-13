package main

import "github.com/gin-gonic/gin"

var userData = map[string]string{
	"johnCena": "WWE champion",
}


// GET
func GetData(c *gin.Context){
	userName := c.Query("name")
	currentTitle, ok := userData[userName]
	if !ok {
		c.JSON(404, gin.H{
			userName: "",
		})
		return
	}

	c.JSON(200, gin.H{
		userName: currentTitle,
	})
}

// POST
func PostData(c *gin.Context){
	userName := c.Query("name")
	currentTitle := c.Query("title")

	if len(userName) == 0 || len(currentTitle) == 0 {
		c.JSON(400, gin.H{
			userName : currentTitle,
		})
		return
	}

	if _,ok := userData[userName];ok {
		c.JSON(409, gin.H{
			"message": "User already exists",
		})
		return
	}

	userData[userName] = currentTitle
	c.JSON(201, gin.H{
		userName: currentTitle,
	})
}

func DeleteData(c *gin.Context){
	userName := c.Query("name")
	currentTitle, ok := userData[userName]
	if !ok {
		c.JSON(400, gin.H{
			userName: "",
		})
		return
	}

	delete(userData,userName)
	c.JSON(200, gin.H{
		userName: currentTitle,
	})
}

func main(){
	r := gin.Default()
	r.GET("/ping",func (c *gin.Context){
		c.JSON(200, gin.H{
			"message":"pong",
		})
	})

	rGroup := r.Group("/data")
	rGroup.GET("", GetData)			// here "/data" is used from r.Group
	rGroup.POST("", PostData)
	rGroup.DELETE("", DeleteData)

	r.Run()  // listen and serve on 0.0.0.0:8080
}