package gui

import (
	"fabricated-calendar/config"
	"fabricated-calendar/internal/database"

	"fyne.io/fyne/v2"
)

type GUI struct {
	App    fyne.App
	Window fyne.Window
	Config config.Config
	User   *database.User
}

func Start(app fyne.App, cfg config.Config) {
	app.Settings().SetTheme(&DefaultTheme{})
	window := app.NewWindow("Fabricated Calendar")

	gui := GUI{
		App:    app,
		Window: window,
		Config: cfg,
	}

	gui.showLogin()

	gui.Window.ShowAndRun()
}
