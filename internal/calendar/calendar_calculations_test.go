package calendar

import (
	"fabricated-calendar/internal/database"
	"testing"
)

func TestCalendarData_DaysPerYear(t *testing.T) {
	tests := []struct {
		name   string
		months []database.Month
		want   int
	}{
		{
			name: "single month",
			months: []database.Month{
				{DaysInMonth: 30},
			},
			want: 30,
		},
		{
			name: "multiple months",
			months: []database.Month{
				{DaysInMonth: 30},
				{DaysInMonth: 31},
				{DaysInMonth: 28},
			},
			want: 89,
		},
		{
			name:   "no months",
			months: nil,
			want:   0,
		},
		{
			name: "zero length month",
			months: []database.Month{
				{DaysInMonth: 0},
				{DaysInMonth: 30},
			},
			want: 30,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := CalendarData{
				Months: tt.months,
			}

			got := data.daysPerYear()

			if got != tt.want {
				t.Errorf(
					"daysPerYear() = %d, want %d",
					got,
					tt.want,
				)
			}
		})
	}
}

func TestCalendarData_MonthStartWeekday(t *testing.T) {
	tests := []struct {
		name       string
		months     []database.Month
		weekdays   []database.Weekday
		monthIndex int
		year       int
		want       int
	}{
		{
			name: "first month year zero",
			months: []database.Month{
				{DaysInMonth: 30},
				{DaysInMonth: 30},
				{DaysInMonth: 30},
			},
			weekdays:   make([]database.Weekday, 7),
			monthIndex: 0,
			year:       0,
			want:       0,
		},
		{
			name: "second month year zero",
			months: []database.Month{
				{DaysInMonth: 30},
				{DaysInMonth: 30},
				{DaysInMonth: 30},
			},
			weekdays:   make([]database.Weekday, 7),
			monthIndex: 1,
			year:       0,
			want:       2,
		},
		{
			name: "third month year zero",
			months: []database.Month{
				{DaysInMonth: 30},
				{DaysInMonth: 30},
				{DaysInMonth: 30},
			},
			weekdays:   make([]database.Weekday, 7),
			monthIndex: 2,
			year:       0,
			want:       4,
		},
		{
			name: "first month year one",
			months: []database.Month{
				{DaysInMonth: 30},
				{DaysInMonth: 30},
				{DaysInMonth: 30},
			},
			weekdays:   make([]database.Weekday, 7),
			monthIndex: 0,
			year:       1,
			want:       6,
		},
		{
			name: "second month year one",
			months: []database.Month{
				{DaysInMonth: 30},
				{DaysInMonth: 30},
				{DaysInMonth: 30},
			},
			weekdays:   make([]database.Weekday, 7),
			monthIndex: 1,
			year:       1,
			want:       1,
		},
		{
			name: "28 day months with five weekdays",
			months: []database.Month{
				{DaysInMonth: 28},
				{DaysInMonth: 28},
			},
			weekdays:   make([]database.Weekday, 5),
			monthIndex: 1,
			year:       0,
			want:       3,
		},
		{
			name: "35 day months with five weekdays",
			months: []database.Month{
				{DaysInMonth: 35},
				{DaysInMonth: 35},
			},
			weekdays:   make([]database.Weekday, 5),
			monthIndex: 1,
			year:       0,
			want:       0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := CalendarData{
				Months:   tt.months,
				Weekdays: tt.weekdays,
			}

			got := data.MonthStartWeekday(
				tt.monthIndex,
				tt.year,
			)

			if got != tt.want {
				t.Errorf(
					"MonthStartWeekday(%d, %d) = %d, want %d",
					tt.monthIndex,
					tt.year,
					got,
					tt.want,
				)
			}
		})
	}
}

func TestCalendarData_MonthStartWeekday_ContinuesAcrossMonths(t *testing.T) {
	data := CalendarData{
		Months: []database.Month{
			{DaysInMonth: 28},
			{DaysInMonth: 28},
			{DaysInMonth: 28},
		},
		Weekdays: make([]database.Weekday, 5),
	}

	tests := []struct {
		monthIndex int
		want       int
	}{
		{
			monthIndex: 0,
			want:       0,
		},
		{
			monthIndex: 1,
			want:       3,
		},
		{
			monthIndex: 2,
			want:       1,
		},
	}

	for _, tt := range tests {
		t.Run("month", func(t *testing.T) {
			got := data.MonthStartWeekday(
				tt.monthIndex,
				0,
			)

			if got != tt.want {
				t.Errorf(
					"MonthStartWeekday(%d, 0) = %d, want %d",
					tt.monthIndex,
					got,
					tt.want,
				)
			}
		})
	}
}

func TestCalendarData_MonthStartWeekday_ContinuesAcrossYears(t *testing.T) {
	data := CalendarData{
		Months: []database.Month{
			{DaysInMonth: 28},
			{DaysInMonth: 28},
		},
		Weekdays: make([]database.Weekday, 5),
	}

	tests := []struct {
		year       int
		monthIndex int
		want       int
	}{
		{
			year:       0,
			monthIndex: 0,
			want:       0,
		},
		{
			year:       0,
			monthIndex: 1,
			want:       3,
		},
		{
			year:       1,
			monthIndex: 0,
			want:       1,
		},
		{
			year:       1,
			monthIndex: 1,
			want:       4,
		},
		{
			year:       2,
			monthIndex: 0,
			want:       2,
		},
	}

	for _, tt := range tests {
		t.Run("calendar", func(t *testing.T) {
			got := data.MonthStartWeekday(
				tt.monthIndex,
				tt.year,
			)

			if got != tt.want {
				t.Errorf(
					"MonthStartWeekday(%d, %d) = %d, want %d",
					tt.monthIndex,
					tt.year,
					got,
					tt.want,
				)
			}
		})
	}
}
