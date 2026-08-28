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

const MainTopDisplay = "main top"

func (g *GUI) mainScreenTopDisplay() fyne.CanvasObject {
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

		g.generateMainScreenMiddleDisplay(MainMiddleDisplay)
		g.generateMainScreenLeftDisplay(MainLeftDisplay)
	})

	// Create Calendar Button
	createCalendarButton := widget.NewButton("Create New Calendar", func() {
		g.generateMainScreenMiddleDisplay(CreateCalendarForm)
	})

	// Edit Calendar Button
	editCalendarButton := widget.NewButton("Edit Calendar", func() {
		g.generateMainScreenMiddleDisplay(EditCalendarForm)
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
		g.showMainScreen()
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
		editCalendarButton,
		deleteCalendarButton,
		layout.NewSpacer(),
		deleteUserButton,
		logoutButton,
	)

	return content
}

func (g *GUI) generateMainScreenTopDisplay(display string) {
	switch display {
	case MainTopDisplay:
		g.TopContainer.Objects = []fyne.CanvasObject{
			g.mainScreenTopDisplay(),
		}

		g.TopContainer.Refresh()
	}
}
