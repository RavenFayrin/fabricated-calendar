package gui

import "fyne.io/fyne/v2/widget"

func (g *GUI) showCalendarSelection() {
	content := widget.NewLabel("test")
	calendarsCard := widget.NewCard("calendar name", "description", content)

	g.Window.SetContent(calendarsCard)
}
