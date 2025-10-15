package models

import (
	"time"
)

//Journal entry is the base struct for all journal entries.

type JournalEntrt struct {
	Content string `json:"content"`
	//Using time.Time for the timestamp field
	Timestamp time.Time `json:"timestamp"`

	//TODO: add the other fields from the python model.
}
