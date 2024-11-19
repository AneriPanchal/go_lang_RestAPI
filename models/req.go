// models/req.go

package models

import "time"

// CreateEventRequest represents the request payload for creating an event.
type CreateEventRequest struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
	Location    string    `json:"location"`
	Capacity    int       `json:"capacity"`
}

// UpdateEventRequest represents the request payload for updating an event.
type UpdateEventRequest struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
	Location    string    `json:"location"`
	Capacity    int       `json:"capacity"`
}
