package gui

import (
	"fabricated-calendar/internal/calendar"
	"fabricated-calendar/internal/database"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/google/uuid"
)

func (g *GUI) showCalendar() {
	top_content := g.topSide()
	left_content := g.leftSide()

	content := container.NewBorder(top_content, nil, left_content, nil, nil)

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
		if g.Calendar == nil || g.Calendar.ID == uuid.Nil {
			g.showError("Calendar not selected.", fmt.Errorf("no calendar selected"))
			return
		}

		err := calendar.DeleteCalendar(g.Config, g.Calendar.ID)
		if err != nil {
			g.showError("Could not delete calendar.", err)
		}

		g.Calendar = &database.Calendar{}
		g.showCalendar()
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
