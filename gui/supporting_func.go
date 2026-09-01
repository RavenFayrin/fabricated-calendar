package gui

import (
	"errors"
	"fabricated-calendar/internal/calendar"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"github.com/google/uuid"
)

func (g *GUI) showError(userMessage string, err error) {
	fyne.LogError(userMessage, err)
	dialog.ShowError(errors.New(userMessage), g.Window)
}

func (g *GUI) checkCalendarSelected() error {
	if g.Calendar == nil || g.Calendar.ID == uuid.Nil {
		return fmt.Errorf("no calendar selected")
	}
	return nil
}

func (g *GUI) fetchCalendarData() error {
	calendarData, err := calendar.GetCalendarData(
		g.Config,
		g.Calendar.ID,
	)
	if err != nil {
		return err
	}
	g.CalendarData = &calendarData
	return nil
}

func (g *GUI) checkCalendarData() error {
	if len(g.CalendarData.Months) == 0 || len(g.CalendarData.Weekdays) == 0 {
		return fmt.Errorf("no months or weekday created")
	}
	return nil
}
