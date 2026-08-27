package calendar

import (
	"context"
	"fabricated-calendar/config"
	"fabricated-calendar/internal/database"
	"strconv"

	"github.com/google/uuid"
)

func CreateWeekday(cfg config.Config, name, order string, calendardID, userId uuid.UUID) error {
	val64, err := strconv.ParseInt(order, 10, 32)
	if err != nil {
		return err
	}

	val32 := int32(val64)

	_, err = cfg.DB.CreateWeekday(context.Background(), database.CreateWeekdayParams{
		Name:       name,
		DayOrder:   val32,
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

func DeleteWeekday(cfg config.Config, weekdayID uuid.UUID) error {
	err := cfg.DB.DeleteWeekday(context.Background(), weekdayID)
	if err != nil {
		return err
	}
	return nil
}
