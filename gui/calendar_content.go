package gui

import (
	"fyne.io/fyne/v2/container"
)

func (g *GUI) showCalendar() {
	topContent := g.topSide()
	g.LeftContainer = container.NewMax(g.leftDisplay())
	g.MiddleContainer = container.NewMax(g.middleDisplay())

	content := container.NewBorder(topContent, nil, g.LeftContainer, nil, g.MiddleContainer)

	g.Window.SetContent(content)
}
