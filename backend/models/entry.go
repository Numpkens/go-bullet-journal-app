package models

import (
	"time"
)

//Journal entry is the base struct for all journal entries.

type JournalEntry struct {
	ID int `json:"id"`
	Content string `json:"content"`
	Type string `json:"type"`
	Date time.Time `json:"timestamp"`
	//TODO: add the other fields from the python model.
}
