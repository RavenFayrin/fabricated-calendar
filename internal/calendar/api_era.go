package calendar

import (
	"context"
	"database/sql"
	"fabricated-calendar/config"
	"fabricated-calendar/internal/database"

	"github.com/google/uuid"
)

func CreateEra(cfg config.Config, name, shorthand, startYear, description string, calendardID, userID uuid.UUID) error {
	valStartYear, err := stringToInt32(startYear)
	if err != nil {
		return err
	}

	var valDesc bool
	var valShort bool

	if description == "" {
		valDesc = false
	} else {
		valDesc = true
	}

	if shorthand == "" {
		valShort = false
	} else {
		valShort = true
	}

	_, err = cfg.DB.CreateEra(context.Background(), database.CreateEraParams{
		Name: name,
		Shorthand: sql.NullString{
			String: shorthand,
			Valid:  valShort},
		StartYear: valStartYear,
		Description: sql.NullString{
			String: description,
			Valid:  valDesc},
		CalendarID: calendardID,
		UserID:     userID,
	})
	if err != nil {
		return err
	}
	return nil
}
