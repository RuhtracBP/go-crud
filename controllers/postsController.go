package controllers

import (
	"github.com/RuhtracBP/go-crud/initializers"
	"github.com/RuhtracBP/go-crud/models"
	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {

	// Get Data of request body

	var body struct {
		Title string
		Body  string
	}
	c.Bind(&body)

	// Create a new post

	//post := models.Post{Title: "Opa", Body: "corpo de Cristo"} //, Birthday: time.Now()

	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil {

		c.JSON(500, gin.H{"message": "Error creating post"})
		return
	}

	// Return the post

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsList(c *gin.Context) {

	// Get The Posts
	var posts []models.Post
	initializers.DB.Find(&posts)

	// Respond Then

	c.JSON(200, gin.H{
		"posts": posts,
	})

}

func PostsShow(c *gin.Context) {

	// Get Id from URL

	id := c.Param("id")

	// Get The Post

	var post models.Post
	initializers.DB.First(&post, id)

	// Respond Then

	c.JSON(200, gin.H{
		"post": post,
	})

}

func PostsUpdate(c *gin.Context) {

	// Get Id from URL

	id := c.Param("id")

	// Get Data of request body

	var body struct {
		Title string
		Body  string
	}
	c.Bind(&body)

	// Get The Post

	var post models.Post
	initializers.DB.First(&post, id)

	// Update The Post

	/* post.Title = body.Title
	post.Body = body.Body

	initializers.DB.Save(&post) */

	initializers.DB.Model(&post).Updates(models.Post{Title: body.Title, Body: body.Body})

	// Respond Then

	c.JSON(200, gin.H{
		"post": post,
	})

}

func PostsDelete(c *gin.Context) {

	// Get Id from URL

	id := c.Param("id")

	// Get The Post

	var post models.Post
	initializers.DB.First(&post, id)

	// Delete The Post

	initializers.DB.Delete(&post)

	// Respond Then

	c.JSON(200, gin.H{
		"Sucesso": "exclusao logica confirmada",
	})

}
