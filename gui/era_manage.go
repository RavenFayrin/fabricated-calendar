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

func (g *GUI) showCreateEra() fyne.CanvasObject {
	eraName := entry(
		EntryOptions{
			PlaceHolder: "Dark Age",
		},
	)

	eraShort := entry(
		EntryOptions{
			PlaceHolder: "DA",
		},
	)

	eraStart := entryNumerical(
		EntryOptions{
			PlaceHolder: "1330",
		},
	)

	eraDesc := entry(
		EntryOptions{
			PlaceHolder: "When knowledge was stolen from the world.",
		},
	)

	form := form(
		[]FormItemOptions{
			{
				Label:  "Era Name",
				Widget: eraName,
			},
			{
				Label:  "Short Era Identifier",
				Widget: eraShort,
			},
			{
				Label:  "Beginning Year",
				Widget: eraStart,
			},
			{
				Label:  "Era Description",
				Widget: eraDesc,
			},
		},
		FormOptions{
			SubmitText: "Create Era",
			OnSubmit: func() {
				err := calendar.CreateEra(
					g.Config,
					eraName.Text,
					eraShort.Text,
					eraStart.Text,
					eraDesc.Text,
					g.Calendar.ID,
					g.User.ID,
				)
				if err != nil {
					g.showError("Unable to create era.", err)
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
			Text: "Create New Era",
		},
	)

	return content
}

func (g *GUI) showEditEra(eraID uuid.UUID) fyne.CanvasObject {
	eraName := entry(
		EntryOptions{
			PlaceHolder: "Dark Age",
		},
	)

	eraShort := entry(
		EntryOptions{
			PlaceHolder: "DA",
		},
	)

	eraStart := entryNumerical(
		EntryOptions{
			PlaceHolder: "1330",
		},
	)

	eraDesc := entry(
		EntryOptions{
			PlaceHolder: "When knowledge was stolen from the world.",
		},
	)

	form := form(
		[]FormItemOptions{
			{
				Label:  "Era Name",
				Widget: eraName,
			},
			{
				Label:  "Short Era Identifier",
				Widget: eraShort,
			},
			{
				Label:  "Beginning Year",
				Widget: eraStart,
			},
			{
				Label:  "Era Description",
				Widget: eraDesc,
			},
		},
		FormOptions{
			SubmitText: "Update Era",
			OnSubmit: func() {
				err := calendar.UpdateEra(
					g.Config,
					eraName.Text,
					eraShort.Text,
					eraStart.Text,
					eraDesc.Text,
					eraID,
				)
				if err != nil {
					g.showError("Unable to update era.", err)
					return
				}

				err = g.fetchCalendarData()
				if err != nil {
					g.showError("Unable to update era.", err)
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
			Text: "Update Era",
		},
	)

	return content
}

func (g *GUI) createEraLables(eras []database.Era) fyne.CanvasObject {
	vbox := container.NewVBox()

	for _, dbEra := range eras {
		eraLabel := textLabel(dbEra.Name)

		editButton := button(
			func() {
				g.generateMainScreenLeftDisplay(EditEraForm, dbEra.ID)
			},
			ButtonOptions{
				Icon: theme.DocumentCreateIcon(),
			},
		)

		deleteButton := button(
			func() {
				err := calendar.DeleteEra(g.Config, dbEra.ID)
				if err != nil {
					g.showError("Could not delete era.", err)
				}
				err = g.fetchCalendarData()
				if err != nil {
					g.showError("Unable to delete era.", err)
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
			eraLabel,
			layout.NewSpacer(),
			editButton,
			deleteButton,
		)

		vbox.Add(row)
	}

	return vbox
}

func (g *GUI) getEras() []database.Era {
	dbEras, err := calendar.GetEras(g.Config, g.Calendar.ID)
	if err != nil {
		g.showError("Unable to get eras.", err)
		return []database.Era{}
	}

	return dbEras
}
