package gui

import (
	"fabricated-calendar/config"
	"fabricated-calendar/internal/auth"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func ShowLogin(window fyne.Window, cfg config.Config) {
	email := widget.NewEntry()
	email.SetPlaceHolder("Email")

	password := widget.NewPasswordEntry()
	password.SetPlaceHolder("Password")

	loginButton := widget.NewButton("Login", func() {
		err := auth.Login(cfg, email.Text, password.Text)

		if err != nil {
			dialog.ShowError(err, window)
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

	window.SetContent(content)
}
