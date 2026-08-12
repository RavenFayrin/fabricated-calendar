package gui

import (
	"fabricated-calendar/internal/auth"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
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

		// Store the logged-in user.
		g.User = &user

		// Actually display the universe selection screen.
		g.Window.SetContent(g.universeScreen())
	})

	createUserButton := widget.NewButton("Create New User", func() {
		g.showUserCreationPopup()
	})

	content := container.NewVBox(
		widget.NewLabel("Fabricated Calendar"),
		username,
		password,
		loginButton,
		createUserButton,
	)

	g.Window.SetContent(
		container.NewCenter(content),
	)
}

// showUserCreationPopup opens the user creation form.
func (g *GUI) showUserCreationPopup() {
	var userDialog *dialog.CustomDialog

	form := g.userCreationForm(
		func() {
			// Submit successfully created the user.
			userDialog.Hide()
		},
		func() {
			// Cancel clears the form and closes the popup.
			userDialog.Hide()
		},
	)

	userDialog = dialog.NewCustom(
		"Create New User",
		"Close",
		form,
		g.Window,
	)

	userDialog.Show()
}

// userCreationForm creates the user creation form.
func (g *GUI) userCreationForm(
	onSubmit func(),
	onCancel func(),
) *fyne.Container {
	username := widget.NewEntry()
	username.SetPlaceHolder("Username")

	password := widget.NewPasswordEntry()
	password.SetPlaceHolder("Password: Must be 8 characters")

	email := widget.NewEntry()
	email.SetPlaceHolder("Email: example@example.com")

	form := &widget.Form{
		Items: []*widget.FormItem{
			{
				Text:   "Username",
				Widget: username,
			},
			{
				Text:   "Password",
				Widget: password,
			},
			{
				Text:   "Email",
				Widget: email,
			},
		},

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

			// Clear the form.
			username.SetText("")
			password.SetText("")
			email.SetText("")

			// Close the popup.
			if onSubmit != nil {
				onSubmit()
			}
		},

		OnCancel: func() {
			// Clear the form.
			username.SetText("")
			password.SetText("")
			email.SetText("")

			// Close the popup.
			if onCancel != nil {
				onCancel()
			}
		},
	}

	return container.NewCenter(form)
}
