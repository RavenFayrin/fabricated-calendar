package calendar

import (
	"context"
	"fabricated-calendar/config"
	"fabricated-calendar/internal/database"

	"github.com/google/uuid"
)

func CreateMonth(cfg config.Config, name string, order, monthLength string, calendardID, userID uuid.UUID) error {
	valOrder, err := stringToInt32(order)
	if err != nil {
		return err
	}

	valLength, err := stringToInt32(monthLength)
	if err != nil {
		return err
	}

	_, err = cfg.DB.CreateMonth(context.Background(), database.CreateMonthParams{
		Name:        name,
		MonthOrder:  valOrder,
		DaysInMonth: valLength,
		CalendarID:  calendardID,
		UserID:      userID,
	})
	if err != nil {
		return err
	}
	return nil
}

func GetMonths(cfg config.Config, calID uuid.UUID) ([]database.Month, error) {
	months, err := cfg.DB.GetMonthsByCalendarId(context.Background(), calID)
	if err != nil {
		return nil, err
	}
	return months, nil
}

func UpdateMonth(cfg config.Config, name string, order, monthLength string, monthID uuid.UUID) error {
	valOrder, err := stringToInt32(order)
	if err != nil {
		return err
	}

	valLength, err := stringToInt32(monthLength)
	if err != nil {
		return err
	}

	_, err = cfg.DB.UpdateMonthById(context.Background(), database.UpdateMonthByIdParams{
		Name:        name,
		MonthOrder:  valOrder,
		DaysInMonth: valLength,
		ID:          monthID,
	})
	if err != nil {
		return err
	}
	return nil
}

func DeleteMonth(cfg config.Config, monthID uuid.UUID) error {
	err := cfg.DB.DeleteMonth(context.Background(), monthID)
	if err != nil {
		return err
	}
	return nil
}
