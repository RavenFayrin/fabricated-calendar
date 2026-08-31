package gui

import (
	"fabricated-calendar/internal/calendar"
	"fabricated-calendar/internal/database"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/google/uuid"

	xwidget "fyne.io/x/fyne/widget"
)

func (g *GUI) showCreateMonth() fyne.CanvasObject {
	monthName := widget.NewEntry()
	monthName.SetPlaceHolder("Month Name")

	monthOrder := xwidget.NewNumericalEntry()
	monthOrder.SetPlaceHolder("Month Order")

	monthLength := xwidget.NewNumericalEntry()
	monthLength.SetPlaceHolder("Month Length")

	submitButton := widget.NewButton("Create Month", func() {
		err := calendar.CreateMonth(
			g.Config,
			monthName.Text,
			monthOrder.Text,
			monthLength.Text,
			g.Calendar.ID,
			g.User.ID,
		)
		if err != nil {
			g.showError("Unable to create month.", err)
			return
		}

		err = g.fetchCalendarData()
		if err != nil {
			g.showError("Unable to update calendar.", err)
			return
		}

		g.generateMainScreenLeftDisplay(MainLeftDisplay)
		g.generateMainScreenMiddleDisplay(MainMiddleDisplay)
	})

	closeButton := widget.NewButton("Close", func() {
		g.generateMainScreenLeftDisplay(MainLeftDisplay)
	})

	content := container.NewVBox(
		widget.NewLabelWithStyle(
			"Create New Month",
			fyne.TextAlignCenter,
			fyne.TextStyle{Bold: true},
		),
		monthName,
		monthOrder,
		monthLength,
		submitButton,
		closeButton,
	)

	return content
}

func (g *GUI) showEditMonth(monthID uuid.UUID) fyne.CanvasObject {
	monthName := widget.NewEntry()
	monthName.SetPlaceHolder("Month Name")

	monthOrder := xwidget.NewNumericalEntry()
	monthOrder.SetPlaceHolder("Month Order")

	monthLength := xwidget.NewNumericalEntry()
	monthLength.SetPlaceHolder("Month Length")

	submitButton := widget.NewButton("Update Month", func() {
		err := calendar.UpdateMonth(
			g.Config,
			monthName.Text,
			monthOrder.Text,
			monthLength.Text,
			monthID,
		)
		if err != nil {
			g.showError("Unable to update month.", err)
			return
		}
		g.generateMainScreenLeftDisplay(MainLeftDisplay)
		g.generateMainScreenMiddleDisplay(MainMiddleDisplay)
	})

	closeButton := widget.NewButton("Close", func() {
		g.generateMainScreenLeftDisplay(MainLeftDisplay)
	})

	content := container.NewVBox(
		widget.NewLabelWithStyle(
			"Update Month",
			fyne.TextAlignCenter,
			fyne.TextStyle{Bold: true},
		),
		monthName,
		monthOrder,
		monthLength,
		submitButton,
		closeButton,
	)

	return content
}

func (g *GUI) createMonthLables(months []database.Month) fyne.CanvasObject {
	vbox := container.NewVBox()

	for _, dbMonth := range months {
		monthLabel := widget.NewLabel(dbMonth.Name)

		editButton := widget.NewButtonWithIcon(
			"",
			theme.DocumentCreateIcon(),
			func() {
				g.generateMainScreenLeftDisplay(EditMonthForm, dbMonth.ID)
			},
		)

		deleteButton := widget.NewButtonWithIcon(
			"",
			theme.DeleteIcon(),
			func() {
				err := calendar.DeleteMonth(g.Config, dbMonth.ID)
				if err != nil {
					g.showError("Could not delete month.", err)
				}
				g.generateMainScreenLeftDisplay(MainLeftDisplay)
			},
		)

		row := container.NewHBox(
			monthLabel,
			layout.NewSpacer(),
			editButton,
			deleteButton,
		)

		vbox.Add(row)
	}

	return vbox
}

func (g *GUI) getMonths() []database.Month {
	dbMonths, err := calendar.GetMonths(g.Config, g.Calendar.ID)
	if err != nil {
		g.showError("Unable to get months.", err)
		return []database.Month{}
	}

	return dbMonths
}
