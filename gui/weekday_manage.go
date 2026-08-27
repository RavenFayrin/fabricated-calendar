package gui

import (
	"fabricated-calendar/internal/calendar"
	"fabricated-calendar/internal/database"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
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

func (g *GUI) getWeekdays() []database.Weekday {
	dbWeekdays, err := calendar.GetWeekdays(g.Config, g.Calendar.ID)
	if err != nil {
		g.showError("Unable to get weekdays.", err)
		return []database.Weekday{}
	}

	return dbWeekdays
}

func (g *GUI) createWeekdayLables(weekdays []database.Weekday) fyne.CanvasObject {
	vbox := container.NewVBox()

	for _, dbWeekday := range weekdays {
		weekdayLabel := widget.NewLabel(dbWeekday.Name)

		editButton := widget.NewButtonWithIcon(
			"",
			theme.DocumentCreateIcon(),
			func() {
				// TODO: edit weekday using dbWeekday.ID
			},
		)

		deleteButton := widget.NewButtonWithIcon(
			"",
			theme.DeleteIcon(),
			func() {
				// TODO: delete weekday using dbWeekday.ID
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
