package gui

import (
	"fabricated-calendar/config"

	"fyne.io/fyne/v2"
)

type GUI struct {
	App    fyne.App
	Window fyne.Window
	Config config.Config
}

func Start(app fyne.App, cfg config.Config) {
	window := app.NewWindow("Fabricated Calendar")

	gui := GUI{
		App:    app,
		Window: window,
		Config: cfg,
	}

	gui.ShowLogin()

	gui.Window.ShowAndRun()
}
