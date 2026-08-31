package gui

import (
	"fabricated-calendar/internal/calendar"
	"fabricated-calendar/internal/database"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/google/uuid"
)

func (g *GUI) showCreateCalendar() fyne.CanvasObject {
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

		g.generateMainScreenTopDisplay(MainTopDisplay)
		g.generateMainScreenMiddleDisplay(MainMiddleDisplay)
		g.generateMainScreenLeftDisplay(MainLeftDisplay)
	})

	closeButton := widget.NewButton("Close", func() {
		g.generateMainScreenMiddleDisplay(MainMiddleDisplay)
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

func (g *GUI) showEditCalendar(calID uuid.UUID) fyne.CanvasObject {
	calName := widget.NewEntry()
	calName.SetPlaceHolder("Calendar Name")

	calDesc := widget.NewEntry()
	calDesc.SetPlaceHolder("Calendar Description")

	submitButton := widget.NewButton("Edit Calendar", func() {
		err := calendar.UpdateCalendar(
			g.Config,
			calName.Text,
			calDesc.Text,
			calID,
		)
		if err != nil {
			g.showError("Unable to edit calendar.", err)
			return
		}

		g.Calendar = &database.Calendar{}
		g.generateMainScreenTopDisplay(MainTopDisplay)
		g.generateMainScreenLeftDisplay(MainLeftDisplay)
		g.generateMainScreenMiddleDisplay(MainMiddleDisplay)
	})

	closeButton := widget.NewButton("Close", func() {
		g.generateMainScreenMiddleDisplay(MainMiddleDisplay)
	})

	content := container.NewPadded(container.NewVBox(
		widget.NewLabelWithStyle(
			"Edit Calendar",
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

func (g *GUI) getCalendars() []database.Calendar {
	dbCalendars, err := calendar.GetCalendars(g.Config, g.User.ID)
	if err != nil {
		g.showError("Unable to get calendars.", err)
		return []database.Calendar{}
	}

	return dbCalendars
}
