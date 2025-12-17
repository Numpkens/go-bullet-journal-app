package models

import (
	"time"
)

// JournalEntry is the base struct for all journal entries.
type JournalEntry struct {
	ID      int       `json:"id"`
	Content string    `json:"content"`
	Type    string    `json:"type"`
	Date    time.Time `json:"date"`
}

