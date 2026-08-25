package gui

import (
	"errors"
	"fabricated-calendar/internal/calendar"
	"fabricated-calendar/internal/database"
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

func (g *GUI) getCalendars() []database.Calendar {
	dbCalendars, err := calendar.GetCalendars(g.Config, g.User.ID)
	if err != nil {
		g.showError("Unable to get calendars.", err)
		return []database.Calendar{}
	}

	return dbCalendars
}
