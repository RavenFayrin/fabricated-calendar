package gui

import (
	"fabricated-calendar/internal/database"

	"fyne.io/fyne/v2"
)

func (g *GUI) makeMenu() *fyne.MainMenu {
	logout := fyne.NewMenuItem(
		"Log Out",
		func() {
			g.User = &database.User{}
			g.Calendar = &database.Calendar{}
			g.showLogin()
		})

	var title string

	err := g.checkUserLoggedIn()
	if err != nil {
		logout.Disabled = true
		title = "Menu"
	} else {
		logout.Disabled = false
		title = g.User.Username
	}

	menu := fyne.NewMenu(
		title,
		logout,
	)

	main := fyne.NewMainMenu(
		menu,
	)

	return main
}
