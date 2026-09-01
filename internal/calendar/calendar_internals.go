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

func (c CalendarData) GetMonthStartWeekday(monthIndex int, year int) int {
	totalDays := 0

	totalDays += c.daysPerYear() * year

	for i := range monthIndex {
		totalDays += int(c.Months[i].DaysInMonth)
	}

	return totalDays % len(c.Weekdays)
}

// func (cs *CalendarSystem) WeeksInYear() int {
// 	return cs.DaysInYear() / len(cs.Weekdays)
// }

// func (cs *CalendarSystem) DateToAbsoluteDay(date *Date) int {

// 	absolute := date.Year * cs.DaysInYear()

// 	for _, m := range cs.Months {
// 		if m.Order < date.Month.Order {
// 			absolute += m.NumDays
// 		}
// 	}

// 	absolute += date.Day

// 	return absolute
// }

// func (cs *CalendarSystem) AbsoluteDayToDate(absolute_day int) (*Date, error) {

// 	daysInYear := cs.DaysInYear()

// 	year := absolute_day / daysInYear
// 	absolute_day %= daysInYear

// 	month := 1

// 	for _, m := range cs.Months {

// 		if absolute_day >= m.NumDays {
// 			absolute_day -= m.NumDays
// 			month++
// 		} else {
// 			break
// 		}
// 	}

// 	return NewDate(cs, year, month, absolute_day)
// }

// func (cs *CalendarSystem) WeekdayName(date *Date) string {

// 	absolute := cs.DateToAbsoluteDay(date)

// 	remainder := (absolute % len(cs.Weekdays)) - 1

// 	if remainder < 0 {
// 		remainder += len(cs.Weekdays)
// 	}

// 	return cs.Weekdays[remainder].Name
// }

// func (cs *CalendarSystem) SortDates(dates []*Date) {
// 	sort.Slice(dates, func(i, j int) bool {
// 		return dates[i].Before(dates[j])
// 	})
// }

// func (cs *CalendarSystem) DaysBetween(a, b *Date) int {
// 	return cs.DateToAbsoluteDay(b) - cs.DateToAbsoluteDay(a)
// }

// func (cs *CalendarSystem) FormattedTimeBetween(a, b *Date) string {

// 	years := 0
// 	months := 0
// 	days := 0

// 	diff := cs.DaysBetween(a, b)

// 	years = diff / cs.DaysInYear()
// 	diff %= cs.DaysInYear()

// 	days += a.Month.NumDays - a.Day
// 	diff -= a.Month.NumDays - a.Day

// 	days += b.Day
// 	diff -= b.Day

// 	for _, m := range cs.Months[a.Month.Order:] {

// 		if diff >= m.NumDays {
// 			diff -= m.NumDays
// 			months++
// 		} else {
// 			break
// 		}
// 	}

// 	return fmt.Sprintf("%d year(s), %d month(s), %d day(s)",
// 		years, months, days)
// }
