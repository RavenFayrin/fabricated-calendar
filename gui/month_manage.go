package gui

import (
	"fabricated-calendar/internal/calendar"
	"fabricated-calendar/internal/database"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"github.com/google/uuid"
)

func (g *GUI) showCreateMonth() fyne.CanvasObject {
	monthName := entry(
		EntryOptions{
			PlaceHolder: "January",
		},
	)

	monthOrder := entryNumerical(
		EntryOptions{
			PlaceHolder: "1",
		},
	)

	monthLength := entryNumerical(
		EntryOptions{
			PlaceHolder: "31",
		},
	)

	form := form(
		[]FormItemOptions{
			{
				Label:  "Month Name",
				Widget: monthName,
			},
			{
				Label:  "Month Order",
				Widget: monthOrder,
			},
			{
				Label:  "Month Length",
				Widget: monthLength,
			},
		},
		FormOptions{
			SubmitText: "Create Month",
			OnSubmit: func() {
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
			},
			CancelText: "Cancel",
			OnCancel: func() {
				g.generateMainScreenLeftDisplay(MainLeftDisplay)
			},
		},
	)

	content := paddedCard(
		form,
		CardOptions{
			Text: "Create New Month",
		},
	)

	return content
}

func (g *GUI) showEditMonth(monthID uuid.UUID) fyne.CanvasObject {
	monthName := entry(
		EntryOptions{
			PlaceHolder: "Febuary",
		},
	)

	monthOrder := entryNumerical(
		EntryOptions{
			PlaceHolder: "2",
		},
	)

	monthLength := entryNumerical(
		EntryOptions{
			PlaceHolder: "28",
		},
	)

	form := form(
		[]FormItemOptions{
			{
				Label:  "Month Name",
				Widget: monthName,
			},
			{
				Label:  "Month Order",
				Widget: monthOrder,
			},
			{
				Label:  "Month Length",
				Widget: monthLength,
			},
		},
		FormOptions{
			SubmitText: "Update Month",
			OnSubmit: func() {
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

				err = g.fetchCalendarData()
				if err != nil {
					g.showError("Unable to update month.", err)
					return
				}

				g.generateMainScreenLeftDisplay(MainLeftDisplay)
				g.generateMainScreenMiddleDisplay(MainMiddleDisplay)
			},
			CancelText: "Cancel",
			OnCancel: func() {
				g.generateMainScreenLeftDisplay(MainLeftDisplay)
			},
		},
	)

	content := paddedCard(
		form,
		CardOptions{
			Text: "Update Month",
		},
	)

	return content
}

func (g *GUI) createMonthLables(months []database.Month) fyne.CanvasObject {
	vbox := container.NewVBox()

	for _, dbMonth := range months {
		monthLabel := textLabel(dbMonth.Name)

		editButton := button(
			func() {
				g.generateMainScreenLeftDisplay(EditMonthForm, dbMonth.ID)
			},
			ButtonOptions{
				Icon: theme.DocumentCreateIcon(),
			},
		)

		deleteButton := button(
			func() {
				err := calendar.DeleteMonth(g.Config, dbMonth.ID)
				if err != nil {
					g.showError("Could not delete month.", err)
				}
				err = g.fetchCalendarData()
				if err != nil {
					g.showError("Unable to delete month.", err)
					return
				}
				g.generateMainScreenLeftDisplay(MainLeftDisplay)
				g.generateMainScreenMiddleDisplay(MainMiddleDisplay)
			},
			ButtonOptions{
				Icon: theme.DeleteIcon(),
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
