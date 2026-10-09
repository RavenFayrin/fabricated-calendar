package gui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"github.com/google/uuid"
)

const MainLeftDisplay = "main left"
const CreateWeekdayForm = "create weekday form"
const EditWeekdayForm = "edit weekday form"
const CreateMonthForm = "create month form"
const EditMonthForm = "edit month form"
const CreateEraForm = "create era form"
const EditEraForm = "edit era form"

func (g *GUI) mainScreenLeftDisplay() fyne.CanvasObject {
	err := g.checkCalendarSelected()
	if err != nil {
		return container.NewWithoutLayout()
	}

	// Create Calendar Parts Buttons
	createWeekdayButton := button(
		func() {
			err := g.checkCalendarSelected()
			if err != nil {
				g.showError("Calendar not selected.", err)
				return
			}

			g.generateMainScreenLeftDisplay(CreateWeekdayForm)
		},
		ButtonOptions{
			Text: "+ Add Weekday",
		},
	)

	createMonthButton := button(
		func() {
			err := g.checkCalendarSelected()
			if err != nil {
				g.showError("Calendar not selected.", err)
				return
			}

			g.generateMainScreenLeftDisplay(CreateMonthForm)
		},
		ButtonOptions{
			Text: "+ Add Month",
		},
	)

	createEraButton := button(
		func() {
			err := g.checkCalendarSelected()
			if err != nil {
				g.showError("Calendar not selected.", err)
				return
			}

			g.generateMainScreenLeftDisplay(CreateEraForm)
		},
		ButtonOptions{
			Text: "+ Add Era",
		},
	)

	// Retrieve Calendar Parts
	dbWeekdays := g.getWeekdays()
	dbMonths := g.getMonths()
	dbEras := g.getEras()

	// Show Calendar Parts
	weekdayLabels := g.createWeekdayLables(dbWeekdays)
	monthLabels := g.createMonthLables(dbMonths)
	eraLabels := g.createEraLables(dbEras)

	// Content
	weekdayContent := container.NewVBox(
		weekdayLabels,
		createWeekdayButton,
	)

	monthContent := container.NewVBox(
		monthLabels,
		createMonthButton,
	)

	eraContent := container.NewVBox(
		eraLabels,
		createEraButton,
	)

	// Calendar tools content
	tools := paddedCard(
		accordion(
			[]AccordionItemOptions{
				{
					Text:    "Weekdays",
					content: weekdayContent,
				},
				{
					Text:    "Months",
					content: monthContent,
				},
				{
					Text:    "Eras",
					content: eraContent,
				},
			},
			AccordionOptions{
				MultiOpen: true,
			},
		),
	)

	// Make the tools scrollable.
	scroll := container.NewVScroll(tools)

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

	case CreateEraForm:
		g.LeftContainer.Objects = []fyne.CanvasObject{
			g.showCreateEra(),
		}

		g.LeftContainer.Refresh()

	case EditEraForm:
		if len(args) == 0 {
			g.showError("Unable to edit era.", fmt.Errorf("era ID was not provided"))
			return
		}
		g.LeftContainer.Objects = []fyne.CanvasObject{
			g.showEditEra(args[0]),
		}

		g.LeftContainer.Refresh()
	}

}
