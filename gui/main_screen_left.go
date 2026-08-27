package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

const MainLeftDisplay = "main left"
const CreateWeekdayForm = "weekday form"
const CreateMonthForm = "month form"

func (g *GUI) mainScreenLeftDisplay() fyne.CanvasObject {
	err := g.checkCalendarSelected()
	if err != nil {
		content := container.NewVBox()
		return content
	}

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
		g.generateMainScreenLeftDisplay(CreateWeekdayForm)
	})

	createMonthButton := widget.NewButton("+ Add Month", func() {
		err := g.checkCalendarSelected()
		if err != nil {
			g.showError("Calendar not selected.", err)
			return
		}
		g.generateMainScreenLeftDisplay(CreateMonthForm)
	})

	// Retrive Calendar Parts
	dbWeekdays := g.getWeekdays()

	// Show Calendar Parts
	weekdayLables := g.createWeekdayLables(dbWeekdays)

	// Edit Calendar Parts

	// Delete Calendar Parts

	content := container.NewVBox(
		titleText,
		weekdaysText,
		weekdayLables,
		createWeekdayButton,
		monthsText,
		createMonthButton,
	)

	return content
}

func (g *GUI) generateMainScreenLeftDisplay(arg string) {
	switch arg {
	case MainLeftDisplay:
		g.LeftContainer.Objects = []fyne.CanvasObject{
			g.mainScreenLeftDisplay(),
		}

		g.LeftContainer.Refresh()
	case CreateWeekdayForm:
		g.LeftContainer.Objects = []fyne.CanvasObject{
			g.showCreateWeekday(),
		}

		g.LeftContainer.Refresh()

	case CreateMonthForm:
		g.LeftContainer.Objects = []fyne.CanvasObject{
			//g.showCreateMonth(),
		}

		g.LeftContainer.Refresh()
	}

}
