package gui

import (
	"fabricated-calendar/internal/calendar"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	xwidget "fyne.io/x/fyne/widget"
)

const MainLeftDisplay = "main"
const CreateWeekdayForm = "weekday form"

func (g *GUI) leftDisplay() fyne.CanvasObject {
	// Text
	titleText := widget.NewLabelWithStyle(
		"Calendar Tools",
		fyne.TextAlignCenter,
		fyne.TextStyle{Bold: true},
	)

	weekdaysText := widget.NewLabelWithStyle(
		"Weekdays",
		fyne.TextAlignCenter,
		fyne.TextStyle{Bold: true},
	)

	monthsText := widget.NewLabelWithStyle(
		"Months",
		fyne.TextAlignCenter,
		fyne.TextStyle{Bold: true},
	)

	// Create Calendar Parts Buttons
	createWeekdayButton := widget.NewButton("+ Add Weekday", func() {
		err := g.checkCalendarSelected()
		if err != nil {
			g.showError("Calendar not selected.", err)
			return
		}
		g.generateLeft(CreateWeekdayForm)
	})

	content := container.NewVBox(
		titleText,
		weekdaysText,
		createWeekdayButton,
		monthsText,
	)

	return content
}

func (g *GUI) showCreateWeekday() fyne.CanvasObject {
	content := g.weekdayCreatationForm()

	return content
}

func (g *GUI) weekdayCreatationForm() *fyne.Container {
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
		g.generateLeft(MainLeftDisplay)
	})

	closeButton := widget.NewButton("Close", func() {
		g.generateLeft(MainLeftDisplay)
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

func (g *GUI) generateLeft(arg string) {
	switch arg {
	case MainLeftDisplay:
		g.LeftContainer.Objects = []fyne.CanvasObject{
			g.leftDisplay(),
		}

		g.LeftContainer.Refresh()
	case CreateWeekdayForm:
		g.LeftContainer.Objects = []fyne.CanvasObject{
			g.showCreateWeekday(),
		}

		g.LeftContainer.Refresh()
	}

}
