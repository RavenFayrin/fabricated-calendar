package gui

import (
	"fabricated-calendar/internal/auth"
	"fabricated-calendar/internal/calendar"
	"fabricated-calendar/internal/database"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func (g *GUI) topSide() fyne.CanvasObject {
	// Get Calendars for Calendar Select
	dbCalendars := g.getCalendars()

	calendarNames := make([]string, 0, len(dbCalendars))
	for i := range dbCalendars {
		calendarNames = append(calendarNames, dbCalendars[i].Name)
	}

	// Calendar Select
	calendarSelect := widget.NewSelect(calendarNames, func(value string) {
		for i := range dbCalendars {
			if dbCalendars[i].Name == value {
				g.Calendar = &dbCalendars[i]
				break
			}
		}

		g.generateMiddle() // Regenerates Middle
		g.generateLeft(MainLeftDisplay)
	})

	// Create Calendar Button
	createCalendarButton := widget.NewButton("Create New Calendar", func() {
		g.showCalendarForm()
	})

	// Delete Calendar Button
	deleteCalendarButton := widget.NewButton("Delete Calendar", func() {
		err := g.checkCalendarSelected()
		if err != nil {
			g.showError("Calendar not selected.", err)
			return
		}

		err = calendar.DeleteCalendar(g.Config, g.Calendar.ID)
		if err != nil {
			g.showError("Could not delete calendar.", err)
		}

		g.Calendar = &database.Calendar{}
		g.showCalendar()
	})

	// Logout Button
	logoutButton := widget.NewButton("Log Out", func() {
		g.User = &database.User{}
		g.Calendar = &database.Calendar{}
		g.showLogin()
	})

	// Delete User Button
	deleteUserButton := widget.NewButton("DELETE USER", func() {
		err := auth.DeleteUser(g.Config, g.User.ID)
		if err != nil {
			g.showError("Could not delete user.", err)
		}
		g.User = &database.User{}
		g.Calendar = &database.Calendar{}
		g.showLogin()
	})

	// Content Creator
	content := container.NewHBox(
		widget.NewLabelWithStyle(
			"Calendar: ",
			fyne.TextAlignCenter,
			fyne.TextStyle{Bold: true},
		),
		calendarSelect,
		createCalendarButton,
		deleteCalendarButton,
		layout.NewSpacer(),
		deleteUserButton,
		logoutButton,
	)

	return content
}

func (g *GUI) getCalendars() []database.Calendar {
	dbCalendars, err := calendar.GetCalendars(g.Config, g.User.ID)
	if err != nil {
		g.showError("Unable to get calendars.", err)
		return []database.Calendar{}
	}

	return dbCalendars
}
