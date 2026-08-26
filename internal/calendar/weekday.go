package calendar

import (
	"context"
	"fabricated-calendar/config"
	"fabricated-calendar/internal/database"

	"github.com/google/uuid"
)

func CreateWeekday(cfg config.Config, name string, order int32, calendardID, weekdayID uuid.UUID) error {
	_, err := cfg.DB.CreateWeekday(context.Background(), database.CreateWeekdayParams{
		Name:       name,
		DayOrder:   order,
		CalendarID: calendardID,
		UserID:     weekdayID,
	})
	if err != nil {
		return err
	}
	return nil
}

// Retrive

func UpdateWeekday(cfg config.Config, name string, order int32, weekdayID uuid.UUID) error {
	_, err := cfg.DB.UpdateWeekdayById(context.Background(), database.UpdateWeekdayByIdParams{
		Name:     name,
		DayOrder: order,
		ID:       weekdayID,
	})
	if err != nil {
		return err
	}
	return nil
}

// Delete
