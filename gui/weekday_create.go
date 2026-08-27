package gui

import (
	"fabricated-calendar/internal/calendar"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	xwidget "fyne.io/x/fyne/widget"
)

func (g *GUI) showCreateWeekday() fyne.CanvasObject {
	weekdayName := widget.NewEntry()
	weekdayName.SetPlaceHolder("Weekday Name")

	weekdayOrder := xwidget.NewNumericalEntry()
	weekdayOrder.SetPlaceHolder("Weekday Order")

	submitButton := widget.NewButton("Create Weekday", func() {
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
		g.generateMainScreenLeftDisplay(MainLeftDisplay)
	})

	closeButton := widget.NewButton("Close", func() {
		g.generateMainScreenLeftDisplay(MainLeftDisplay)
	})

	content := container.NewVBox(
		widget.NewLabelWithStyle(
			"Create New Weekday",
			fyne.TextAlignCenter,
			fyne.TextStyle{Bold: true},
		),
		weekdayName,
		weekdayOrder,
		submitButton,
		closeButton,
	)

	return content
}
