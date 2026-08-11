package gui

import (
	"fabricated-calendar/internal/auth"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func (g *GUI) ShowLogin() {
	username := widget.NewEntry()
	username.SetPlaceHolder("Username")

	password := widget.NewPasswordEntry()
	password.SetPlaceHolder("Password")

	loginButton := widget.NewButton("Login", func() {
		match, err := auth.Login(g.Config, username.Text, password.Text)
		if err != nil {
			g.showError("Unable to login. Please check email and password.", err)
			return
		}
		if match == true {
			g.showCalendarSelection()
		}
	})

	createUserButton := widget.NewButton("Create New User", func() {
		g.ShowCreateUser()
	})

	content := container.NewVBox(
		widget.NewLabel("Fabricated Calendar"),
		username,
		password,
		loginButton,
		createUserButton,
	)

	g.Window.SetContent(content)
}
