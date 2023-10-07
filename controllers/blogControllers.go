package controllers

import (
	database "App/database"
	models "App/models"
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var db *sql.DB

func PostBlog(c *gin.Context) {
	// Initialize the database connection
	db, err := database.InitDB()
	if err != nil {
		log.Printf("Error initializing the database: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	defer db.Close()

	// Bind the JSON data to the newPost struct
	var newPost models.Article
	if err := c.ShouldBindJSON(&newPost); err != nil {
		log.Printf("Error in binding post data: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Define the SQL query for inserting a new blog post
	query := `INSERT INTO blogPosts (title, content) VALUES (?, ?)`

	// Prepare and execute the SQL statement
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Printf("Error preparing SQL statement: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(newPost.Title, newPost.Content)
	if err != nil {
		log.Printf("Error inserting new blog post: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Blog added successfully"})
}
