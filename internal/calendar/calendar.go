package calendar

import (
	"context"
	"database/sql"
	"fabricated-calendar/config"
	"fabricated-calendar/internal/database"

	"github.com/google/uuid"
)

func CreateCalendar(cfg config.Config, name, desc string, userID uuid.UUID) error {
	var val bool

	if desc == "" {
		val = false
	} else {
		val = true
	}

	_, err := cfg.DB.CreateCalendar(context.Background(), database.CreateCalendarParams{
		Name: name,
		Description: sql.NullString{
			String: desc,
			Valid:  val},
		UserID: userID,
	})
	if err != nil {
		return err
	}

	return nil
}

func GetCalendars(cfg config.Config, userID uuid.UUID) ([]database.Calendar, error) {
	calendars, err := cfg.DB.GetCalendarsByUserId(context.Background(), userID)
	if err != nil {
		return nil, err
	}
	return calendars, nil
}
