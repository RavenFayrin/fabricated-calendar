package gui

import (
	"fabricated-calendar/internal/calendar"
	"fabricated-calendar/internal/database"
	"fmt"

	"fyne.io/fyne/v2"
	"github.com/google/uuid"
)

func (g *GUI) showCreateCalendar() fyne.CanvasObject {
	calName := entry()

	calDesc := entry(
		EntryOptions{
			MultiLine: true,
		},
	)

	form := form(
		[]FormItemOptions{
			{
				Label:  "Calendar Name",
				Widget: calName,
			},
			{
				Label:  "Calendar Description",
				Widget: calDesc,
			},
		},
		FormOptions{
			SubmitText: "Create Calendar",
			OnSubmit: func() {
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
			},
			CancelText: "Cancel",
			OnCancel: func() {
				g.generateMainScreenMiddleDisplay(MainMiddleDisplay)
			},
		},
	)

	content := paddedCard(
		form,
		CardOptions{
			Text: "Create New Calendar",
		},
	)

	return content
}

func (g *GUI) showEditCalendar(calID uuid.UUID) fyne.CanvasObject {
	calName := entry(
		EntryOptions{
			PlaceHolder: "Change Calendar Name",
		},
	)

	calDesc := entry(
		EntryOptions{
			PlaceHolder: "Change Calendar Description",
			MultiLine:   true,
		},
	)

	form := form(
		[]FormItemOptions{
			{
				Label:  "Calendar Name",
				Widget: calName,
			},
			{
				Label:  "Calendar Description",
				Widget: calDesc,
			},
		},
		FormOptions{
			SubmitText: "Update Calendar",
			OnSubmit: func() {
				err := calendar.UpdateCalendar(
					g.Config,
					calName.Text,
					calDesc.Text,
					calID,
				)
				if err != nil {
					g.showError("Unable to update calendar.", err)
					return
				}

				g.Calendar = &database.Calendar{}
				g.generateMainScreenTopDisplay(MainTopDisplay)
				g.generateMainScreenLeftDisplay(MainLeftDisplay)
				g.generateMainScreenMiddleDisplay(MainMiddleDisplay)
			},
			CancelText: "Cancel",
			OnCancel: func() {
				g.generateMainScreenMiddleDisplay(MainMiddleDisplay)
			},
		},
	)

	content := paddedCard(
		form,
		CardOptions{
			Text: fmt.Sprintf("Update %s", g.Calendar.Name),
		},
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
