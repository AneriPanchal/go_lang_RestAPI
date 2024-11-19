package main

import (
	"database/sql"
	//"database/db"
	"eventapp/models"
	"eventapp/config"
	"eventapp/controller"
	"eventapp/manager"
	"eventapp/services"
	"github.com/labstack/echo/v4"
	"log"
	"os"
	"fmt"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	if err := godotenv.Load(".config"); err != nil {
		log.Println("No .config file found")
	}

	// Initialize configuration
	cfg := config.LoadConfig()

	//postgres................
	postgresConn := "postgres://postgres:admin@localhost:5433/postgres?sslmode=disable"

	// Connect to PostgreSQL server
	serverDB, err := sql.Open("postgres", postgresConn)
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL server: %v", err)
	}
	defer serverDB.Close()

	// Create the `eventdb` database if it doesn't exist
	if err := models.CreateDatabase(serverDB); err != nil {
		log.Fatalf("Failed to create database: %v", err)
	}

	// Connect to the `eventdb` database
	eventDBConn := "postgres://postgres:admin@localhost:5433/eventdb?sslmode=disable"
	eventDB, err := sql.Open("postgres", eventDBConn)
	if err != nil {
		log.Fatalf("Failed to connect to eventdb: %v", err)
	}
	defer eventDB.Close()

	// Create the `events` table
	if err := models.CreateEventsTable(eventDB); err != nil {
		log.Fatalf("Failed to create events table: %v", err)
	}

	fmt.Println("Database and table setup completed successfully.")


	
	// Initialize dependencies
	eventManager := &manager.EventManager{DB: cfg.DB}
	eventService := &services.EventService{EventManager: eventManager}
	eventController := &controllers.EventController{EventService: eventService}

	// Set up Echo
	e := echo.New()

	// Routes
	e.POST("/events", eventController.CreateEvent)
	e.GET("/events", eventController.GetAllEvents)
	e.PUT("/events/:id", eventController.UpdateEvent)
    e.DELETE("/events/:id", eventController.DeleteEvent)


	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	e.Logger.Fatal(e.Start(":" + port))
}







