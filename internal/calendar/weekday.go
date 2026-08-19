package calendar

import (
	"context"
	"fabricated-calendar/config"
	"fabricated-calendar/internal/database"

	"github.com/google/uuid"
)

func CreateWeekday(cfg config.Config, name string, order int32, calendardID, userID uuid.UUID) error {
	_, err := cfg.DB.CreateWeekday(context.Background(), database.CreateWeekdayParams{
		Name:       name,
		DayOrder:   order,
		CalendarID: calendardID,
		UserID:     userID,
	})
	if err != nil {
		return err
	}
	return nil
}
