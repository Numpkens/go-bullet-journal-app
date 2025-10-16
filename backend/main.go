package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//Handler for the /entries route
//

func getEntries(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to GoJo Your Online Bullet Journal",
		"entries": []string{},
	})
}

func main() {
	router := gin.Default()
	router.GET("/entries", getEntries)
	// Run the server on port 8000. It will listen for incoming requests here.
	router.Run("localhost:8000")
}
