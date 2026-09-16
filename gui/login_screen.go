package gui

import (
	"fabricated-calendar/internal/auth"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func (g *GUI) showLogin() {
	username := entry()

	password := entry(
		EntryOptions{
			Password: true,
		},
	)

	form := form(
		[]FormItemOptions{
			{
				Label:  "Username",
				Widget: username,
			},
			{
				Label:  "Password",
				Widget: password,
			},
		},
		FormOptions{
			SubmitText: "Login",
			OnSubmit: func() {
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
			},
		},
	)

	createUserButton := button(
		func() {
			g.showUserCreation()
		},
		ButtonOptions{
			Text: "Create New User",
		},
	)

	content := paddedCard(
		container.NewVBox(
			form,
			createUserButton,
		),
		CardOptions{
			Text: "Fabricated Calendar Login",
		},
	)

	g.Window.SetContent(content)
}

func (g *GUI) showUserCreation() {
	content := g.userCreationForm()

	g.Window.SetContent(content)
}

func (g *GUI) userCreationForm() *fyne.Container {
	username := entry()

	password := entry(
		EntryOptions{
			Password: true,
		},
	)

	email := entry(
		EntryOptions{
			PlaceHolder: "example@example.com",
		},
	)

	form := form(
		[]FormItemOptions{
			{
				Label:  "Username",
				Widget: username,
			},
			{
				Label:  "Password",
				Widget: password,
			},
			{
				Label:  "Email",
				Widget: email,
			},
		},
		FormOptions{
			SubmitText: "Create User",
			OnSubmit: func() {
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
			},
			CancelText: "Cancel",
			OnCancel: func() {
				g.showLogin()
			},
		},
	)

	content := paddedCard(
		form,
		CardOptions{
			Text: "Create New User",
		},
	)

	return content
}
