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
	dbMonths := g.getMonths()

	// Show Calendar Parts
	weekdayLables := g.createWeekdayLables(dbWeekdays)
	monthLables := g.createMonthLables(dbMonths)

	content := container.NewVBox(
		titleText,
		weekdaysText,
		weekdayLables,
		createWeekdayButton,
		monthsText,
		monthLables,
		createMonthButton,
	)

	return content
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
