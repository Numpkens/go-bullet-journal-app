package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// write temporary hardcoded data
var journalEntries = []models.JournalEntry{
	{ID: 1, Content: "Draft first API endpoint", Type: "Task", Timestamp: time.Now()},
	{ID: 2, Content: "Yulia Class", Type: "Event", Timestamp: time.Now().Add(time.Hour * 2)},
	{ID: 3, Content: "Having a wonderful day", Type: "Journal", Timestamp: time.Now().Add(time.Hour * 2)},
}
//Handler for the /entries route
func getEntries(c *gin.Context) {
	c.JSON(http.StatusOK, journalEntries)
}

func main() {
	router := gin.Default()
	router.GET("/entries", getEntries)
	// Run the server on port 8000. It will listen for incoming requests here.
	router.Run("localhost:8000")
}
