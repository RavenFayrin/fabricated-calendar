package config

import "fabricated-calendar/internal/database"

type Config struct {
	DB *database.Queries
}
