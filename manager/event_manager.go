package manager

import (
	"context"
	"database/sql"
	"eventapp/models"
)

type EventManager struct {
	DB *sql.DB
}

func (m *EventManager) CreateEvent(ctx context.Context, event models.Event) (string, error) {
	query := `
		INSERT INTO events (title, description, date, location)
		VALUES ($1, $2, $3, $4)
		RETURNING id`
	var id string
	err := m.DB.QueryRowContext(ctx, query, event.Title, event.Description, event.Date, event.Location).Scan(&id)
	return id, err
}

func (m *EventManager) GetAllEvents(ctx context.Context) ([]models.Event, error) {
	query := `SELECT id, title, description, date, location FROM events`
	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []models.Event
	for rows.Next() {
		var event models.Event
		err := rows.Scan(&event.ID, &event.Title, &event.Description, &event.Date, &event.Location)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, rows.Err()
}


func (m *EventManager) UpdateEvent(ctx context.Context, id string, event models.Event) error {
	query := `
		UPDATE events
		SET title = $1, description = $2, date = $3, location = $4
		WHERE id = $5
	`
	_, err := m.DB.ExecContext(ctx, query, event.Title, event.Description, event.Date, event.Location, id)
	return err
}
func (m *EventManager) DeleteEvent(ctx context.Context, id string) error {
	query := `
		DELETE FROM events
		WHERE id = $1
	`
	_, err := m.DB.ExecContext(ctx, query, id)
	return err
}
