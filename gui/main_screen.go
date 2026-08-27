package gui

import (
	"fyne.io/fyne/v2/container"
)

func (g *GUI) showMainScreen() {
	topContent := g.mainScreenTopDisplay()
	g.LeftContainer = container.NewMax(g.mainScreenLeftDisplay())
	g.MiddleContainer = container.NewMax(g.mainScreenMiddleCalendarDisplay())

	content := container.NewBorder(topContent, nil, g.LeftContainer, nil, g.MiddleContainer)

	g.Window.SetContent(content)
}
