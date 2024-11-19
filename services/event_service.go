
package services

import (
	"context"
	"eventapp/manager"
	"eventapp/models"
)

type EventService struct {
	EventManager *manager.EventManager
}

func (s *EventService) CreateEvent(ctx context.Context, event models.Event) (string, error) {
	return s.EventManager.CreateEvent(ctx, event)
}
func (s *EventService) GetAllEvents(ctx context.Context) ([]models.Event, error) {
	return s.EventManager.GetAllEvents(ctx)
}
func (s *EventService) UpdateEvent(ctx context.Context, id string, event models.Event) error {
	return s.EventManager.UpdateEvent(ctx, id, event)
}
func (s *EventService) DeleteEvent(ctx context.Context, id string) error {
	return s.EventManager.DeleteEvent(ctx, id)
}
