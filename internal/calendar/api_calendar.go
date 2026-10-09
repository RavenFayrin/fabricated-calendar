package calendar

import (
	"context"
	"database/sql"
	"fabricated-calendar/config"
	"fabricated-calendar/internal/database"

	"github.com/google/uuid"
)

func CreateCalendar(cfg config.Config, name, desc string, userID uuid.UUID) error {
	val := nullStringIdentifier(desc)

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

func UpdateCalendar(cfg config.Config, name, desc string, calendarID uuid.UUID) error {
	val := nullStringIdentifier(desc)

	_, err := cfg.DB.UpdateCalendarByID(context.Background(), database.UpdateCalendarByIDParams{
		Name: name,
		Description: sql.NullString{
			String: desc,
			Valid:  val},
		ID: calendarID,
	})
	if err != nil {
		return err
	}

	return nil
}

func DeleteCalendar(cfg config.Config, id uuid.UUID) error {
	err := cfg.DB.DeleteCalendar(context.Background(), id)
	if err != nil {
		return err
	}
	return nil
}
