package gui

import (
	"fabricated-calendar/internal/calendar"
	"fabricated-calendar/internal/database"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
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
				err := g.fetchCalendarData()
				if err != nil {
					g.showError("Unable to select calendar.", err)
				}
				g.DisplayMonthIndex = 0
				g.DisplayYear = 0
				break
			}
		}

		g.generateMainScreenMiddleDisplay(MainMiddleDisplay)
		g.generateMainScreenLeftDisplay(MainLeftDisplay)
	})
	calendarSelect.PlaceHolder = "Select Calendar"

	// Create Calendar Button
	createCalendarButton := button(
		func() {
			g.generateMainScreenMiddleDisplay(CreateCalendarForm)
		},
		ButtonOptions{
			Text: "Create New Calendar",
			Icon: theme.ContentAddIcon(),
		},
	)

	// Edit Calendar Button
	editCalendarButton := button(
		func() {
			g.generateMainScreenMiddleDisplay(EditCalendarForm)
		},
		ButtonOptions{
			Text: "Edit Calendar",
			Icon: theme.DocumentCreateIcon(),
		},
	)

	// Delete Calendar Button
	deleteCalendarButton := button(
		func() {
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
		},
		ButtonOptions{
			Text: "Delete Calendar",
			Icon: theme.DeleteIcon(),
		},
	)

	calendarCard := widget.NewCard(
		"",
		"",
		container.NewGridWithColumns(
			4,
			calendarSelect,
			editCalendarButton,
			deleteCalendarButton,
			createCalendarButton,
		),
	)

	// Content Creator
	content := container.NewPadded(
		calendarCard,
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
