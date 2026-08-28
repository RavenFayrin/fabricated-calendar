package calendar

import (
	"context"
	"fabricated-calendar/config"
	"fabricated-calendar/internal/database"

	"github.com/google/uuid"
)

func CreateWeekday(cfg config.Config, name, order string, calendardID, userId uuid.UUID) error {
	val, err := stringToInt32(order)
	if err != nil {
		return err
	}

	_, err = cfg.DB.CreateWeekday(context.Background(), database.CreateWeekdayParams{
		Name:       name,
		DayOrder:   val,
		CalendarID: calendardID,
		UserID:     userId,
	})
	if err != nil {
		return err
	}
	return nil
}

func GetWeekdays(cfg config.Config, calID uuid.UUID) ([]database.Weekday, error) {
	weekdays, err := cfg.DB.GetWeekdaysByCalendarId(context.Background(), calID)
	if err != nil {
		return nil, err
	}
	return weekdays, nil
}

func UpdateWeekday(cfg config.Config, name, order string, weekdayID uuid.UUID) error {
	val, err := stringToInt32(order)
	if err != nil {
		return err
	}

	_, err = cfg.DB.UpdateWeekdayById(context.Background(), database.UpdateWeekdayByIdParams{
		Name:     name,
		DayOrder: val,
		ID:       weekdayID,
	})
	if err != nil {
		return err
	}
	return nil
}

func DeleteWeekday(cfg config.Config, weekdayID uuid.UUID) error {
	err := cfg.DB.DeleteWeekday(context.Background(), weekdayID)
	if err != nil {
		return err
	}
	return nil
}
