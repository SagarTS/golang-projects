package controllers

import (
	"gin-crud/models"

	"gin-crud/initializers"

	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context){
	// get data of request body
	var body struct {
		Title string
		Body string
	}

	c.Bind(&body)

	// create a post
	post := models.Post{Title: body.Title, Body:body.Body}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return 
	}

	// return result
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context){
	// get all the data
	var posts []models.Post
	initializers.DB.Find(&posts)

	// respond to request
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostsShow(c *gin.Context){
	// Get id from URL
	id := c.Param("id")

	// Get data
	var post models.Post
	initializers.DB.First(&post, id)

	// respond
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsUpdate(c *gin.Context){
	// Get id from url
	id := c.Param("id")

	// Get data off request body
	var body struct{
		Title string
		Body string
	}
	c.Bind(&body)

	// Find the post we are updating
	var post models.Post
	initializers.DB.First(&post,id)
	
	// Update
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body: body.Body,
	})

	//Response updated data
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsDelete(c *gin.Context){
	// Get id from url
	id:= c.Param("id")
			
	// Delete it
	initializers.DB.Delete(&models.Post{}, id)

	//Respond
	c.Status(200)
}