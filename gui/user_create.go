package gui

import (
	"fabricated-calendar/internal/auth"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func (g *GUI) ShowCreateUser() {
	username := widget.NewEntry()
	username.SetPlaceHolder("Username")

	password := widget.NewPasswordEntry()
	password.SetPlaceHolder("Password: Must be 8 characters")

	email := widget.NewEntry()
	email.SetPlaceHolder("Email: example@example.com")

	createButton := widget.NewButton("Create", func() {
		err := auth.CreateUser(g.Config, username.Text, password.Text, email.Text)
		if err != nil {
			g.showError("Unable to create user.", err)
			return
		}
	})

	cancelButton := widget.NewButton("Cancel", func() {
		g.ShowLogin()
	})

	content := container.NewVBox(
		widget.NewLabel("Fabricated Calendar"),
		username,
		password,
		email,
		cancelButton,
		createButton,
	)

	g.Window.SetContent(content)
}
