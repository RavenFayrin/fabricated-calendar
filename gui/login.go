package gui

import (
	"fabricated-calendar/internal/auth"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func (g *GUI) ShowLogin() {
	email := widget.NewEntry()
	email.SetPlaceHolder("Email")

	password := widget.NewPasswordEntry()
	password.SetPlaceHolder("Password")

	loginButton := widget.NewButton("Login", func() {
		err := auth.Login(g.Config, email.Text, password.Text)

		if err != nil {
			g.showError("Unable to login. Please check email and password.", err)
			return
		}
	})

	createUserButton := widget.NewButton("Create New User", func() {
		//ShowCreateUser(window)
	})

	content := container.NewVBox(
		widget.NewLabel("Fabricated Calendar"),
		email,
		password,
		loginButton,
		createUserButton,
	)

	g.Window.SetContent(content)
}
