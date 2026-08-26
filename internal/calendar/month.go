package calendar

import (
	"context"
	"fabricated-calendar/config"
	"fabricated-calendar/internal/database"

	"github.com/google/uuid"
)

func CreateMonth(cfg config.Config, name string, order, numDays int32, calendardID, userID uuid.UUID) error {
	_, err := cfg.DB.CreateMonth(context.Background(), database.CreateMonthParams{
		Name:        name,
		MonthOrder:  order,
		DaysInMonth: numDays,
		CalendarID:  calendardID,
		UserID:      userID,
	})
	if err != nil {
		return err
	}
	return nil
}

// retrive

func UpdateMonth(cfg config.Config, name string, order, numDays int32, monthID uuid.UUID) error {
	_, err := cfg.DB.UpdateMonthById(context.Background(), database.UpdateMonthByIdParams{
		Name:        name,
		MonthOrder:  order,
		DaysInMonth: numDays,
		ID:          monthID,
	})
	if err != nil {
		return err
	}
	return nil
}

func DeteleMonth(cfg config.Config, monthID uuid.UUID) error {
	err := cfg.DB.DeleteMonth(context.Background(), monthID)
	if err != nil {
		return err
	}
	return nil
}
