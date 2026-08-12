package gui

import (
	"fabricated-calendar/internal/calendar"
	"fabricated-calendar/internal/database"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func (g *GUI) universeScreen() fyne.CanvasObject {
	grid := g.universeGrid()

	return container.NewVBox(
		widget.NewLabelWithStyle(
			"Your Universes",
			fyne.TextAlignCenter,
			fyne.TextStyle{Bold: true},
		),
		grid,
	)
}

// universeGrid creates the grid containing all of the user's universes.
func (g *GUI) universeGrid() *fyne.Container {
	grid := container.NewGridWrap(
		fyne.NewSize(250, 180),
	)

	g.refreshUniverseGrid(grid)

	return grid
}

// refreshUniverseGrid gets the user's universes from the database
// and rebuilds the grid.
func (g *GUI) refreshUniverseGrid(grid *fyne.Container) {
	universes, err := calendar.UniversesGetByUserId(g.Config, *g.User)
	if err != nil {
		g.showError("Unable to fetch universes.", err)
		return
	}

	// Remove everything currently in the grid.
	grid.Objects = nil

	// Add a card for each universe.
	for _, universe := range universes {
		grid.Add(g.universeCard(universe))
	}

	// Create New Universe is always the last card.
	grid.Add(g.createUniverseCard(grid))

	grid.Refresh()
}

// universeCard creates a visual card for an existing universe.
func (g *GUI) universeCard(universe database.Universe) fyne.CanvasObject {
	name := widget.NewLabelWithStyle(
		universe.Name,
		fyne.TextAlignCenter,
		fyne.TextStyle{Bold: true},
	)

	description := widget.NewLabel(universe.Description.String)
	description.Wrapping = fyne.TextWrapWord

	// TODO: Make the entire card clickable when you are ready
	// to open a universe.
	card := container.NewBorder(
		name,
		nil,
		nil,
		nil,
		description,
	)

	return container.NewPadded(card)
}

// createUniverseCard creates the final card in the grid.
func (g *GUI) createUniverseCard(grid *fyne.Container) fyne.CanvasObject {
	return widget.NewButton("Create New Universe", func() {
		g.showUniverseCreationPopup(grid)
	})
}

// showUniverseCreationPopup opens the universe creation form.
func (g *GUI) showUniverseCreationPopup(grid *fyne.Container) {
	form := g.universeCreationForm(func() {
		// Refresh the grid after successfully creating a universe.
		g.refreshUniverseGrid(grid)
	})

	dialog.ShowCustom(
		"Create New Universe",
		"Close",
		form,
		g.Window,
	)
}

// universeCreationForm creates the form used to create a new universe.
//
// onCreated is called after the universe has successfully been created.
func (g *GUI) universeCreationForm(onCreated func()) *fyne.Container {
	universeName := widget.NewEntry()
	universeName.SetPlaceHolder("Universe name")

	universeDesc := widget.NewMultiLineEntry()
	universeDesc.SetPlaceHolder("Universe description")

	form := &widget.Form{
		Items: []*widget.FormItem{
			{
				Text:   "Universe Name",
				Widget: universeName,
			},
			{
				Text:   "Universe Description",
				Widget: universeDesc,
			},
		},

		OnSubmit: func() {
			name := strings.TrimSpace(universeName.Text)
			desc := strings.TrimSpace(universeDesc.Text)

			_, err := calendar.UniverseCreate(
				g.Config,
				*g.User,
				name,
				desc,
			)
			if err != nil {
				g.showError("Unable to create universe.", err)
				return
			}

			// Tell the universe grid to refresh.
			if onCreated != nil {
				onCreated()
			}

			// Clear the form.
			universeName.SetText("")
			universeDesc.SetText("")
		},

		OnCancel: func() {
			universeName.SetText("")
			universeDesc.SetText("")
		},
	}

	return container.NewCenter(form)
}
