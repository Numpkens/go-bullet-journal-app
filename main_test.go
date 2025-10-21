package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"module github.com/Numpkens/go-bullet-journal-app"

// Helper function to set up the Gin router for testing
func setupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/entries", getEntries)
	router.POST("/entries", postEntry) // This is the new handler we'll need to create
	return router
}

// Test function for the POST /entries route
func TestPostEntry(t *testing.T) {
	// 1. Arrange: Prepare the test data
	router := setupRouter()

	newEntry := models.JournalEntry{
		Content: "A brand new entry for testing!",
		Type:    "Test",
	}

	// Convert the Go struct to a JSON payload
	jsonValue, _ := json.Marshal(newEntry)

	// 2. Act: Execute the request
	// Create a new HTTP request
	req, _ := http.NewRequest("POST", "/entries", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	// Record the response
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// 3. Assert: Verify the results

	// Check the status code
	if w.Code != http.StatusCreated {
		t.Errorf("Expected status code %d, got %d", http.StatusCreated, w.Code)
	}

	// Check if the entry was added to the slice (check the new length)
	expectedLength := 4 // We start with 3 entries, so we expect 4
	if len(journalEntries) != expectedLength {
		t.Errorf("Expected journalEntries length %d, got %d", expectedLength, len(journalEntries))
	}
}
