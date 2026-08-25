package gui

import (
	"fyne.io/fyne/v2/container"
)

func (g *GUI) showCalendar() {
	top_content := g.topSide()
	left_content := g.leftSide()
	g.CalendarContainer = container.NewMax(g.middleDisplay())

	content := container.NewBorder(top_content, nil, left_content, nil, g.CalendarContainer)

	g.Window.SetContent(content)
}
