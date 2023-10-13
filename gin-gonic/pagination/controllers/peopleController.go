package controllers

import (
	"net/http"
	"strconv"

	"go-pagination/helper"
	"go-pagination/initializers"
	"go-pagination/models"

	"github.com/gin-gonic/gin"
)



func PeopleIndexGET(c *gin.Context) {
	// Get page number
	perPage := 10
	page := 1
	pageStr := c.Param("page")

	if pageStr != ""{
		page,_ = strconv.Atoi(pageStr)
	}

	pagination := helper.GetPaginationData(page, perPage, models.Person{}, "/people")

	// Get the people
	var people []models.Person
	initializers.DB.Limit(10).Offset(pagination.Offset).Find(&people)

	// Render the page
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"people": people,
		"pagination": pagination,
	})
}