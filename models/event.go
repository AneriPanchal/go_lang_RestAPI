package models

import ("time"
	   "database/sql"
	   "fmt"
)
type Event struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
	Location    string    `json:"location"`
}
func CreateDatabase(db *sql.DB) error {
	query := `
	CREATE DATABASE eventdb TEMPLATE template0;
	`
	_, err := db.Exec(query)
	if err != nil {
		if err.Error() == "pq: database \"eventdb\" already exists" {
			fmt.Println("Database already exists, skipping creation.")
			return nil
		}
		return fmt.Errorf("error creating database: %w", err)
	}
	fmt.Println("Database 'eventdb2' created successfully.")
	return nil
}
func CreateEventsTable(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS events (
		id SERIAL PRIMARY KEY,
		title VARCHAR(255) NOT NULL,
		description TEXT,
		date TIMESTAMP NOT NULL,
		location VARCHAR(255)
	);
	`
	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("error creating events table: %w", err)
	}
	fmt.Println("Table 'events' created successfully.")
	return nil
}


