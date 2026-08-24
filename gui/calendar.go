package gui

import (
	"fabricated-calendar/internal/calendar"
	"fabricated-calendar/internal/database"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func (g *GUI) showCalendar() {
	top_content := g.topSide()

	content := container.NewBorder(top_content, nil, nil, nil, nil)

	g.Window.SetContent(content)
}

func (g *GUI) topSide() fyne.CanvasObject {
	dbCalendars := g.getCalendars()

	calendarNames := make([]string, 0, len(dbCalendars))
	for _, cal := range dbCalendars {
		calendarNames = append(calendarNames, cal.Name)
	}

	calendarSelect := widget.NewSelect(calendarNames, func(value string) {
		for _, cal := range dbCalendars {
			if cal.Name == value {
				g.Calendar = &cal
				break
			}
		}
	})

	createCalendarButton := widget.NewButton("Create New Calendar", func() {
		g.showCalendarForm()
	})

	createLogoutButton := widget.NewButton("Log Out", func() {
		g.User = &database.User{}
		g.Calendar = &database.Calendar{}
		g.showLogin()
	})

	content := container.NewHBox(
		calendarSelect,
		createCalendarButton,
		createLogoutButton,
	)

	return content
}

func (g *GUI) leftSide() {

}

func (g *GUI) getCalendars() []database.Calendar {
	dbCalendars, err := calendar.GetCalendars(g.Config, g.User.ID)
	if err != nil {
		g.showError("Unable to get calendars.", err)
		return []database.Calendar{}
	}

	return dbCalendars
}

func (g *GUI) showCalendarForm() {
	content := g.calendarCreatationForm()

	g.Window.SetContent(content)
}

func (g *GUI) calendarCreatationForm() *fyne.Container {
	calName := widget.NewEntry()
	calName.SetPlaceHolder("Calendar Name")

	calDesc := widget.NewEntry()
	calDesc.SetPlaceHolder("Calendar Description")

	submitButton := widget.NewButton("Create Calendar", func() {
		err := calendar.CreateCalendar(
			g.Config,
			calName.Text,
			calDesc.Text,
			g.User.ID,
		)
		if err != nil {
			g.showError("Unable to create calendar.", err)
			return
		}

		g.showCalendar()
	})

	closeButton := widget.NewButton("Close", func() {
		g.showCalendar()
	})

	content := container.NewPadded(container.NewVBox(
		widget.NewLabelWithStyle(
			"Create New Calendar",
			fyne.TextAlignCenter,
			fyne.TextStyle{Bold: true},
		),
		calName,
		calDesc,
		submitButton,
		closeButton,
	))

	return content
}
