package controllers

import (
	"eventapp/models"
	"eventapp/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

type EventController struct {
	EventService *services.EventService
}

func (c *EventController) CreateEvent(ctx echo.Context) error {
	var event models.Event
	if err := ctx.Bind(&event); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	id, err := c.EventService.CreateEvent(ctx.Request().Context(), event)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create event"})
	}

	return ctx.JSON(http.StatusCreated, map[string]string{"id": id})
}

func (c *EventController) GetAllEvents(ctx echo.Context) error {
	events, err := c.EventService.GetAllEvents(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch events"})
	}

	return ctx.JSON(http.StatusOK, events)
}



func (c *EventController) UpdateEvent(ctx echo.Context) error {
	// Parse the event ID from the URL parameter
	id := ctx.Param("id")
	if id == "" {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Event ID is required"})
	}

	// Bind the request body to the Event model
	var event models.Event
	if err := ctx.Bind(&event); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	// Update the event using the service
	if err := c.EventService.UpdateEvent(ctx.Request().Context(), id, event); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update event"})
	}

	return ctx.JSON(http.StatusOK, map[string]string{"message": "Event updated successfully"})
}
func (c *EventController) DeleteEvent(ctx echo.Context) error {
	// Parse the event ID from the URL parameter
	id := ctx.Param("id")
	if id == "" {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Event ID is required"})
	}

	// Delete the event using the service
	if err := c.EventService.DeleteEvent(ctx.Request().Context(), id); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete event"})
	}

	return ctx.JSON(http.StatusOK, map[string]string{"message": "Event deleted successfully"})
}
