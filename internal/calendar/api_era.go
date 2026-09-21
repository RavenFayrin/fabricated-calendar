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

	valDesc := nullStringIdentifier(description)
	valShort := nullStringIdentifier(shorthand)

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

func GetEras(cfg config.Config, calID uuid.UUID) ([]database.Era, error) {
	eras, err := cfg.DB.GetErasByCalendarId(context.Background(), calID)
	if err != nil {
		return nil, err
	}
	return eras, nil
}

func UpdateEra(cfg config.Config, name, shorthand, startYear, description string, eraID uuid.UUID) error {
	valStartYear, err := stringToInt32(startYear)
	if err != nil {
		return err
	}

	valDesc := nullStringIdentifier(description)
	valShort := nullStringIdentifier(shorthand)

	_, err = cfg.DB.UpdateEraById(context.Background(), database.UpdateEraByIdParams{
		Name: name,
		Shorthand: sql.NullString{
			String: shorthand,
			Valid:  valShort},
		StartYear: valStartYear,
		Description: sql.NullString{
			String: description,
			Valid:  valDesc},
		ID: eraID,
	})
	if err != nil {
		return err
	}

	return nil
}

func DeleteEra(cfg config.Config, eraID uuid.UUID) error {
	err := cfg.DB.DeleteEra(context.Background(), eraID)
	if err != nil {
		return err
	}
	return nil
}
