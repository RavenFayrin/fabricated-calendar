package gui

import (
	"fabricated-calendar/config"
	"fabricated-calendar/gui/theme"
	"fabricated-calendar/internal/database"

	"fyne.io/fyne/v2"
)

type GUI struct {
	App             fyne.App
	Window          fyne.Window
	Config          config.Config
	User            *database.User
	Calendar        *database.Calendar
	TopContainer    *fyne.Container
	LeftContainer   *fyne.Container
	MiddleContainer *fyne.Container
}

func Start(app fyne.App, cfg config.Config) {
	app.Settings().SetTheme(theme.DefaultTheme{})
	window := app.NewWindow("Fabricated Calendar")

	gui := GUI{
		App:    app,
		Window: window,
		Config: cfg,
	}

	gui.showLogin()

	gui.Window.ShowAndRun()
}
