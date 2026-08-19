package gui

import (
	"fabricated-calendar/internal/calendar"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

//import "fyne.io/fyne/v2/container"

func (g *GUI) showCalendar() {
	top_content := g.topSide()
	//left_content := //
	//middle_content := //

	content := container.NewBorder(top_content, nil, nil, nil, nil)

	g.Window.SetContent(content)
}

func (g *GUI) topSide() fyne.CanvasObject {
	createCalendarButton := widget.NewButton("Create New Calendar", func() {
		g.showCalendarForm()
	})

	content := container.NewHBox(
		createCalendarButton,
	)

	return content
}

func (g *GUI) leftSide() {

}

func (g *GUI) showCalendarForm() {
	content := g.calendarCreatationForm()

	g.Window.SetContent(content)
}

func (g *GUI) calendarCreatationForm() *fyne.Container {
	calName := widget.NewEntry()
	calName.SetPlaceHolder("Calendar Name")

	calDesc := widget.NewEntry()
	calDesc.SetPlaceHolder("Calendar Description")

	submitButton := widget.NewButton("Create Calendar", func() {
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

		g.showCalendar()
	})

	closeButton := widget.NewButton("Close", func() {
		g.showCalendar()
	})

	content := container.NewPadded(container.NewVBox(
		widget.NewLabelWithStyle(
			"Create New Calendar",
			fyne.TextAlignCenter,
			fyne.TextStyle{Bold: true},
		),
		calName,
		calDesc,
		submitButton,
		closeButton,
	))

	return content
}
