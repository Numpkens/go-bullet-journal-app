package main

import (
	"github.com/gin-gonic/gin"
	"go-bullet-journal-app/backend/models"
	"net/http"
	"time"
)

// Global in-memory data store
var journalEntries = []models.JournalEntry{
	{ID: 1, Content: "Prioritize DSA/Math for 2 hours.", Type: "Task", Date: time.Now()},
	{ID: 2, Content: "ESL Teaching commitment.", Type: "Event", Date: time.Now().Add(time.Hour * 5)},
	{ID: 3, Content: "Configuration fixed!", Type: "Journal", Date: time.Now()},
}

// Handler for the /entries route
func getEntries(c *gin.Context) {
	c.JSON(http.StatusOK, journalEntries)
}

func main() {
	router := gin.Default()
	router.GET("/entries", getEntries)
	router.Run("localhost:8000")
}
