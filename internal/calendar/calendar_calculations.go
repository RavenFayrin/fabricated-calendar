package calendar

import (
	"context"
	"fabricated-calendar/config"
	"fabricated-calendar/internal/database"

	"github.com/google/uuid"
)

type CalendarData struct {
	Calendar database.Calendar
	Weekdays []database.Weekday
	Months   []database.Month
}

func GetCalendarData(
	cfg config.Config,
	calendarID uuid.UUID,
) (CalendarData, error) {
	var data CalendarData

	cal, err := cfg.DB.GetCalendarById(context.Background(), calendarID)
	if err != nil {
		return data, err
	}

	weekdays, err := cfg.DB.GetWeekdaysByCalendarId(
		context.Background(),
		calendarID,
	)
	if err != nil {
		return data, err
	}

	months, err := cfg.DB.GetMonthsByCalendarId(
		context.Background(),
		calendarID,
	)
	if err != nil {
		return data, err
	}

	data.Calendar = cal
	data.Weekdays = weekdays
	data.Months = months

	return data, nil
}

func (c CalendarData) daysPerYear() int {
	total := 0

	for _, month := range c.Months {
		total += int(month.DaysInMonth)
	}
	return total
}

func (c CalendarData) MonthStartWeekday(monthIndex int, year int) int {
	totalDays := 0

	totalDays += c.daysPerYear() * year

	for i := range monthIndex {
		totalDays += int(c.Months[i].DaysInMonth)
	}

	return totalDays % len(c.Weekdays)
}
