package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func (g *GUI) middleDisplay() fyne.CanvasObject {
	err := g.checkCalendarSelected()
	if err != nil {
		content := container.NewVBox(
			widget.NewLabelWithStyle(
				"No Calendar Selected",
				fyne.TextAlignCenter,
				fyne.TextStyle{Bold: true},
			),
			widget.NewLabelWithStyle(
				"Select a calendar above to begin.",
				fyne.TextAlignCenter,
				fyne.TextStyle{Bold: false},
			),
		)
		return content
	}
	content := container.NewVBox(
		widget.NewLabelWithStyle(
			g.Calendar.Name,
			fyne.TextAlignCenter,
			fyne.TextStyle{Bold: true},
		),
		widget.NewLabelWithStyle(
			"Month - Year",
			fyne.TextAlignCenter,
			fyne.TextStyle{Bold: true},
		),
	)

	return content
}

func (g *GUI) generateCalendar() {
	g.CalendarContainer.Objects = []fyne.CanvasObject{
		g.middleDisplay(),
	}

	g.CalendarContainer.Refresh()
}
