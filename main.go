package main

import (
	"database/sql"
	"log"
	"os"

	"fyne.io/fyne/v2/app"
	"github.com/joho/godotenv"

	"fabricated-calendar/config"
	"fabricated-calendar/gui"
	"fabricated-calendar/internal/database"
)

func main() {
	godotenv.Load()
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL must be set")
	}

	dbConn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Error opening database: %s", err)
	}
	dbQueries := database.New(dbConn)

	cfg := config.Config{
		DB: dbQueries,
	}

	app := app.New()

	gui.Start(app, cfg)
}
