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
	calName := entry(
		EntryOptions{
			PlaceHolder: "Calendar Name",
		},
	)

	calDesc := entry(
		EntryOptions{
			PlaceHolder: "Calendar Description",
		},
	)

	form := widget.NewForm(
		widget.NewFormItem("Calendar Name", calName),
		widget.NewFormItem("Calendar Description", calDesc),
	)

	form.OnSubmit = func() {
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
	}

	form.OnCancel = func() {
		g.generateMainScreenMiddleDisplay(MainMiddleDisplay)
	}

	content := container.NewPadded(
		widget.NewCard(
			"Create New Calendar",
			"",
			form,
		),
	)

	return content
}

func (g *GUI) showEditCalendar(calID uuid.UUID) fyne.CanvasObject {
	calName := entry(
		EntryOptions{
			PlaceHolder: "Calendar Name",
		},
	)

	calDesc := entry(
		EntryOptions{
			PlaceHolder: "Calendar Description",
		},
	)

	submitButton := button(
		func() {
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
		},
		ButtonOptions{
			Text: "Edit Calendar",
		},
	)

	closeButton := button(
		func() {
			g.generateMainScreenMiddleDisplay(MainMiddleDisplay)
		},
		ButtonOptions{
			Text: "Close",
		},
	)

	content := container.NewPadded(
		container.NewVBox(
			textLabel(
				"Edit Calendar",
				LabelOptions{
					Alignment: fyne.TextAlignCenter,
					Bold:      true,
				},
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
