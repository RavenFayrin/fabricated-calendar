package gui

import (
	"fabricated-calendar/config"

	"fyne.io/fyne/v2"
)

func Start(app fyne.App, cfg config.Config) {
	window := app.NewWindow("Fabricated Calendar")

	ShowLogin(window, cfg)

	window.ShowAndRun()
}
