package gui

import (
	"fabricated-calendar/internal/auth"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func (g *GUI) showLogin() {
	username := widget.NewEntry()
	username.SetPlaceHolder("Username")

	password := widget.NewPasswordEntry()
	password.SetPlaceHolder("Password")

	loginButton := widget.NewButton("Login", func() {
		user, err := auth.Login(
			g.Config,
			username.Text,
			password.Text,
		)
		if err != nil {
			g.showError(
				"Unable to login. Please check username and password.",
				err,
			)
			return
		}

		g.User = &user

		g.showMainScreen()
	})

	createUserButton := widget.NewButton("Create New User", func() {
		g.showUserCreation()
	})

	content := container.NewPadded(container.NewVBox(
		widget.NewLabelWithStyle(
			"Fabricated Calendar",
			fyne.TextAlignCenter,
			fyne.TextStyle{Bold: true},
		),
		username,
		password,
		loginButton,
		createUserButton,
	))

	g.Window.SetContent(content)
}

func (g *GUI) showUserCreation() {
	content := g.userCreationForm()

	g.Window.SetContent(content)
}

func (g *GUI) userCreationForm() *fyne.Container {
	username := widget.NewEntry()
	username.SetPlaceHolder("Username")

	password := widget.NewPasswordEntry()
	password.SetPlaceHolder("Password: Must be 8 characters")

	email := widget.NewEntry()
	email.SetPlaceHolder("Email: example@example.com")

	submitButton := widget.NewButton("Create User", func() {
		err := auth.CreateUser(
			g.Config,
			username.Text,
			password.Text,
			email.Text,
		)
		if err != nil {
			g.showError("Unable to create user.", err)
			return
		}

		g.showLogin()
	})

	closeButton := widget.NewButton("Close", func() {
		g.showLogin()
	})

	content := container.NewVBox(
		widget.NewLabelWithStyle(
			"Create New User",
			fyne.TextAlignCenter,
			fyne.TextStyle{Bold: true},
		),
		username,
		password,
		email,
		submitButton,
		closeButton,
	)

	return content
}
