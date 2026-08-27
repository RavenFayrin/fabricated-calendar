package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

const MainLeftDisplay = "main left"
const CreateWeekdayForm = "weekday form"

func (g *GUI) mainScreenLeftDisplay() fyne.CanvasObject {
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

	content := container.NewVBox(
		titleText,
		weekdaysText,
		createWeekdayButton,
		monthsText,
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
	}

}
