package model

import (
	"time"
)

type Priority int

const (
	Low Priority = iota
	Medium
	High
)

type Task struct {
	ID          int
	Title       string
	Description string
	Priority    Priority
	Completed   bool
	CreatedAt   time.Time
}

// Returning enum number of string priority
func ParsePriority(priority string) *Priority {
	switch priority {
	case "medium":
		medium := Medium
		return &medium
	case "high":
		high := High
		return &high
	}

	low := Low
	return &low
}

// Parsing priority enum to string value
func PriorityToString(priority Priority) string {
	switch priority {
	case Medium:
		return "medium"
	case High:
		return "high"
	}

	return "low"
}
