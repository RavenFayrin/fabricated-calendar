package gui

import (
	"fabricated-calendar/internal/auth"
	"fabricated-calendar/internal/database"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func (g *GUI) makeMenu() *fyne.MainMenu {
	// Menu Action Buttons
	logout := fyne.NewMenuItemWithIcon(
		"Log Out",
		theme.LogoutIcon(),
		func() {
			g.User = &database.User{}
			g.Calendar = &database.Calendar{}
			g.showLogin()
		})

	deleteUser := fyne.NewMenuItemWithIcon(
		"Delete User",
		theme.DeleteIcon(),
		func() {
			dialog.NewCustomConfirm(
				"Delete User",
				"Confirm",
				"Cancel",
				widget.NewLabel("Are you sure you want to delete this account? This action cannot be undone."),
				func(confirmed bool) {
					if !confirmed {
						return
					}
					err := auth.DeleteUser(g.Config, g.User.ID)
					if err != nil {
						g.showError("Could not delete user.", err)
						return
					}
					g.User = &database.User{}
					g.Calendar = &database.Calendar{}
					g.showLogin()
				},
				g.Window,
			).Show()
		},
	)

	// Menu Items Display
	manageUser := fyne.NewMenuItemWithIcon(
		"Manage Account",
		theme.AccountIcon(),
		nil,
	)
	manageUser.ChildMenu = fyne.NewMenu(
		"",
		deleteUser,
	)

	// User Error Check
	var title string

	err := g.checkUserLoggedIn()
	if err != nil {
		logout.Disabled = true
		manageUser.Disabled = true
		title = "Menu"
	} else {
		logout.Disabled = false
		logout.Disabled = false
		title = g.User.Username
	}

	menu := fyne.NewMenu(
		title,
		manageUser,
		logout,
	)

	main := fyne.NewMainMenu(
		menu,
	)

	return main
}
