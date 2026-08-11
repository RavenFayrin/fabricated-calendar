package main

import (
	"database/sql"
	"log"
	"os"

	"fyne.io/fyne/v2/app"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"fabricated-calendar/config"
	"fabricated-calendar/gui"
	"fabricated-calendar/internal/database"
)

func main() {
	godotenv.Load()
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL must be set")
	}

	dbConn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Error opening database: %s", err)
	}
	dbQueries := database.New(dbConn)

	cfg := config.Config{
		DB: dbQueries,
	}

	app := app.NewWithID("com.ravenfayrin.fabricatedcalendar")

	gui.Start(app, cfg)
}
