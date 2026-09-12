package gui

import (
	"fabricated-calendar/internal/calendar"
	"fabricated-calendar/internal/database"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"github.com/google/uuid"
)

func (g *GUI) showCreateWeekday() fyne.CanvasObject {
	weekdayName := entry(
		EntryOptions{
			PlaceHolder: "Weekday Name",
		},
	)

	weekdayOrder := entryNumerical(
		EntryOptions{
			PlaceHolder: "Weekday Order",
		},
	)

	submitButton := button(
		func() {
			err := calendar.CreateWeekday(
				g.Config,
				weekdayName.Text,
				weekdayOrder.Text,
				g.Calendar.ID,
				g.User.ID,
			)
			if err != nil {
				g.showError("Unable to create weekday.", err)
				return
			}

			err = g.fetchCalendarData()
			if err != nil {
				g.showError("Unable to update calendar.", err)
				return
			}

			g.generateMainScreenLeftDisplay(MainLeftDisplay)
			g.generateMainScreenMiddleDisplay(MainMiddleDisplay)
		},
		ButtonOptions{
			Text: "Create Weekday",
		},
	)

	closeButton := button(
		func() {
			g.generateMainScreenLeftDisplay(MainLeftDisplay)
		},
		ButtonOptions{
			Text: "Close",
		},
	)

	content := container.NewPadded(
		container.NewVBox(
			textLabel(
				"Create New Weekday",
				LabelOptions{
					Alignment: fyne.TextAlignCenter,
					Bold:      true,
				},
			),
			weekdayName,
			weekdayOrder,
			submitButton,
			closeButton,
		))

	return content
}

func (g *GUI) showEditWeekday(weekdayID uuid.UUID) fyne.CanvasObject {
	weekdayName := entry(
		EntryOptions{
			PlaceHolder: "Weekday Name",
		},
	)

	weekdayOrder := entryNumerical(
		EntryOptions{
			PlaceHolder: "Weekday Order",
		},
	)

	submitButton := button(
		func() {
			err := calendar.UpdateWeekday(
				g.Config,
				weekdayName.Text,
				weekdayOrder.Text,
				weekdayID,
			)
			if err != nil {
				g.showError("Unable to update weekday.", err)
				return
			}

			err = g.fetchCalendarData()
			if err != nil {
				g.showError("Unable to update weekday.", err)
				return
			}

			g.generateMainScreenLeftDisplay(MainLeftDisplay)
			g.generateMainScreenMiddleDisplay(MainMiddleDisplay)
		},
		ButtonOptions{
			Text: "Update Weekday",
		},
	)

	closeButton := button(
		func() {
			g.generateMainScreenLeftDisplay(MainLeftDisplay)
		},
		ButtonOptions{
			Text: "Close",
		},
	)

	content := container.NewPadded(
		container.NewVBox(
			textLabel(
				"Update Weekday",
				LabelOptions{
					Alignment: fyne.TextAlignCenter,
					Bold:      true,
				},
			),
			weekdayName,
			weekdayOrder,
			submitButton,
			closeButton,
		))

	return content
}

func (g *GUI) createWeekdayLables(weekdays []database.Weekday) fyne.CanvasObject {
	vbox := container.NewVBox()

	for _, dbWeekday := range weekdays {
		weekdayLabel := textLabel(dbWeekday.Name)

		editButton := button(
			func() {
				g.generateMainScreenLeftDisplay(EditWeekdayForm, dbWeekday.ID)
			},
			ButtonOptions{
				Icon: theme.DocumentCreateIcon(),
			},
		)

		deleteButton := button(
			func() {
				err := calendar.DeleteWeekday(g.Config, dbWeekday.ID)
				if err != nil {
					g.showError("Could not delete weekday.", err)
				}

				err = g.fetchCalendarData()
				if err != nil {
					g.showError("Unable to delete weekday.", err)
					return
				}

				g.generateMainScreenLeftDisplay(MainLeftDisplay)
				g.generateMainScreenMiddleDisplay(MainMiddleDisplay)
			},
			ButtonOptions{
				Icon: theme.DeleteIcon(),
			},
		)

		row := container.NewHBox(
			weekdayLabel,
			layout.NewSpacer(),
			editButton,
			deleteButton,
		)

		vbox.Add(row)
	}

	return vbox
}

func (g *GUI) getWeekdays() []database.Weekday {
	dbWeekdays, err := calendar.GetWeekdays(g.Config, g.Calendar.ID)
	if err != nil {
		g.showError("Unable to get weekdays.", err)
		return []database.Weekday{}
	}

	return dbWeekdays
}
