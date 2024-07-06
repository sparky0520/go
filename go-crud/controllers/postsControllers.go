package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/sparky0520/go-crud/initializers"
	"github.com/sparky0520/go-crud/models"
)

func PostsCreate(c *gin.Context) {
	// Get data off request body
	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)

	// Create a post
	post := models.Post{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&post) // pass pointer of data to Create
	if result.Error != nil {
		c.Status(400)
		return
	}

	//Return it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {
	// Get all posts from the database
	var posts []models.Post
	initializers.DB.Find(&posts)

	// Return them
	c.JSON(200, gin.H{
		"posts": posts,
	})

}

func PostsShow(c *gin.Context) {
	// Get id passed with url
	id := c.Param("id")

	// Find and get the post from the database
	var post models.Post
	initializers.DB.First(&post, id)

	// Return the post
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsUpdate(c *gin.Context) {
	// Get id off the url
	id := c.Param("id")

	// Get the data off the request body
	var body struct {
		Title string
		Body  string
	}
	c.Bind(&body)

	// Update the database with req data for given id
	var post models.Post
	initializers.DB.First(&post, id)

	initializers.DB.Model(&post).
		Updates(models.Post{
			Title: body.Title,
			Body:  body.Body,
		})

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsDelete(c *gin.Context) {
	// Get id off the url
	id := c.Param("id")

	// Find and Delete the post
	initializers.DB.Delete(&models.Post{}, id)

	// Return the post deleted
	c.Status(200)
}
