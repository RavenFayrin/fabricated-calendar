package gui

import (
	"fabricated-calendar/config"
	"fabricated-calendar/gui/theme"
	"fabricated-calendar/internal/calendar"
	"fabricated-calendar/internal/database"

	"fyne.io/fyne/v2"
)

type GUI struct {
	App    fyne.App
	Window fyne.Window

	Config   config.Config
	User     *database.User
	Calendar *database.Calendar

	CalendarData *calendar.CalendarData

	DisplayYear       int32
	DisplayMonthIndex int32

	TopContainer    *fyne.Container
	LeftContainer   *fyne.Container
	MiddleContainer *fyne.Container
}

func Start(app fyne.App, cfg config.Config) {
	app.Settings().SetTheme(theme.WoodlandTheme{})
	window := app.NewWindow("Fabricated Calendar")

	gui := GUI{
		App:    app,
		Window: window,
		Config: cfg,
	}

	window.SetMainMenu(gui.makeMenu())
	window.SetMaster()

	gui.showLogin()

	gui.Window.ShowAndRun()
}

func (g *GUI) makeMenu() *fyne.MainMenu {
	logout := fyne.NewMenuItem("Log Out", func() {
		g.User = &database.User{}
		g.Calendar = &database.Calendar{}
		g.showLogin()
	})
	err := g.checkUserLoggedIn()
	if err != nil {
		logout.Disabled = true
	} else {
		logout.Disabled = false
	}

	menu := fyne.NewMenu(
		"Menu",
		logout)
	main := fyne.NewMainMenu(
		menu,
	)
	return main
}
