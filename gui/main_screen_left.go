package gui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/google/uuid"
)

const MainLeftDisplay = "main left"
const CreateWeekdayForm = "create weekday form"
const EditWeekdayForm = "edit weekday form"
const CreateMonthForm = "create month form"
const EditMonthForm = "edit month form"

func (g *GUI) mainScreenLeftDisplay() fyne.CanvasObject {
	err := g.checkCalendarSelected()
	if err != nil {
		return container.NewVBox()
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

	// Retrieve Calendar Parts
	dbWeekdays := g.getWeekdays()
	dbMonths := g.getMonths()

	// Show Calendar Parts
	weekdayLabels := g.createWeekdayLables(dbWeekdays)
	monthLabels := g.createMonthLables(dbMonths)

	// Calendar tools content
	tools := container.NewVBox(
		titleText,
		weekdaysText,
		weekdayLabels,
		createWeekdayButton,
		monthsText,
		monthLabels,
		createMonthButton,
	)

	// Add padding between the content and the left side
	// of the window.
	padded := container.NewPadded(tools)

	// Make the tools scrollable.
	scroll := container.NewVScroll(padded)

	return scroll
}

func (g *GUI) generateMainScreenLeftDisplay(display string, args ...uuid.UUID) {
	switch display {
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

	case EditWeekdayForm:
		if len(args) == 0 {
			g.showError("Unable to edit weekday.", fmt.Errorf("weekday ID was not provided"))
			return
		}
		g.LeftContainer.Objects = []fyne.CanvasObject{
			g.showEditWeekday(args[0]),
		}

		g.LeftContainer.Refresh()

	case CreateMonthForm:
		g.LeftContainer.Objects = []fyne.CanvasObject{
			g.showCreateMonth(),
		}

		g.LeftContainer.Refresh()

	case EditMonthForm:
		if len(args) == 0 {
			g.showError("Unable to edit month.", fmt.Errorf("month ID was not provided"))
			return
		}
		g.LeftContainer.Objects = []fyne.CanvasObject{
			g.showEditMonth(args[0]),
		}

		g.LeftContainer.Refresh()
	}

}
