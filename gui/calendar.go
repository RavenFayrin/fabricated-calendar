package gui

import (
	"fabricated-calendar/internal/calendar"
	"fabricated-calendar/internal/database"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func (g *GUI) showCalendarWindow() {
	top_content := g.topSide()
	left_content := g.leftSide()
	middle_content := g.middleDisplay()

	content := container.NewBorder(top_content, nil, left_content, nil, middle_content)

	g.Window.SetContent(content)
}

func (g *GUI) topSide() fyne.CanvasObject {
	dbCalendars := g.getCalendars()

	calendarNames := make([]string, 0, len(dbCalendars))
	for i := range dbCalendars {
		calendarNames = append(calendarNames, dbCalendars[i].Name)
	}

	calendarSelect := widget.NewSelect(calendarNames, func(value string) {
		for i := range dbCalendars {
			if dbCalendars[i].Name == value {
				g.Calendar = &dbCalendars[i]
				break
			}
		}
	})

	createCalendarButton := widget.NewButton("Create New Calendar", func() {
		g.showCalendarForm()
	})

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
		g.showCalendarWindow()
	})

	logoutButton := widget.NewButton("Log Out", func() {
		g.User = &database.User{}
		g.Calendar = &database.Calendar{}
		g.showLogin()
	})

	content := container.NewHBox(
		widget.NewLabelWithStyle(
			"Calendar: ",
			fyne.TextAlignCenter,
			fyne.TextStyle{Bold: true},
		),
		calendarSelect,
		createCalendarButton,
		deleteCalendarButton,
		logoutButton,
	)

	return content
}

func (g *GUI) leftSide() fyne.CanvasObject {
	title := widget.NewLabelWithStyle(
		"Calendar Tools",
		fyne.TextAlignCenter,
		fyne.TextStyle{Bold: true},
	)

	content := container.NewVBox(
		title,
	)

	return content
}

func (g *GUI) middleDisplay() fyne.CanvasObject {
	err := g.checkCalendarSelected()
	if err != nil {
		content := container.NewVBox(
			widget.NewLabelWithStyle(
				"No Calendar Selected",
				fyne.TextAlignCenter,
				fyne.TextStyle{Bold: true},
			),
			widget.NewLabelWithStyle(
				"Select a calendar above to begin.",
				fyne.TextAlignCenter,
				fyne.TextStyle{Bold: false},
			),
		)
		return content
	}
	content := container.NewVBox(
		widget.NewLabelWithStyle(
			g.Calendar.Name,
			fyne.TextAlignCenter,
			fyne.TextStyle{Bold: true},
		),
		widget.NewLabelWithStyle(
			"Month - Year",
			fyne.TextAlignCenter,
			fyne.TextStyle{Bold: true},
		),
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
