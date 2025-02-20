package main

import (
	"os"

	"github.com/Lakasha-hub/micro-yoga-mp.git/internal/platform/db"
	"github.com/gin-gonic/gin"
)

func main() {

	// Connect to database
	db.Database()

	// Create a new Gin router
	router := gin.Default()

	// Define the route
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Microservicio Activooo",
		})
	})

	router.Run(":%s", os.Getenv("PORT"))
}
