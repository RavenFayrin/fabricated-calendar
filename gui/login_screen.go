package gui

import (
	"fabricated-calendar/internal/auth"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func (g *GUI) showLogin() {
	username := entry(
		EntryOptions{
			PlaceHolder: "Enter Username",
		},
	)

	password := entry(
		EntryOptions{
			PlaceHolder: "Enter Password",
			Password:    true,
		},
	)

	form := widget.NewForm(
		widget.NewFormItem("Username", username),
		widget.NewFormItem("Password", password),
	)

	form.OnSubmit = func() {
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

		g.Window.SetMainMenu(g.makeMenu())
		g.showMainScreen()
	}
	form.SubmitText = "Login"

	createUserButton := button(
		func() {
			g.showUserCreation()
		},
		ButtonOptions{
			Text: "Create New User",
		},
	)

	content := container.NewPadded(
		widget.NewCard(
			"Fabricated Calendar Login",
			"",
			container.NewVBox(
				form,
				createUserButton,
			),
		))

	g.Window.SetContent(content)
}

func (g *GUI) showUserCreation() {
	content := g.userCreationForm()

	g.Window.SetContent(content)
}

func (g *GUI) userCreationForm() *fyne.Container {
	username := entry(
		EntryOptions{
			PlaceHolder: "Username",
		},
	)

	password := entry(
		EntryOptions{
			PlaceHolder: "Password",
			Password:    true,
		},
	)

	email := entry(
		EntryOptions{
			PlaceHolder: "example@example.com",
		},
	)

	form := widget.NewForm(
		widget.NewFormItem(
			"Username",
			username,
		),
		widget.NewFormItem(
			"Password",
			password,
		),
		widget.NewFormItem(
			"Email",
			email,
		),
	)

	form.OnSubmit = func() {
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
	}
	form.SubmitText = "Create User"

	form.OnCancel = func() {
		g.showLogin()
	}
	form.CancelText = "Cancel"

	content := container.NewPadded(
		widget.NewCard(
			"Create New User",
			"",
			form,
		))

	return content
}
