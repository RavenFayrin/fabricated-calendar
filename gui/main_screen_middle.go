package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

const MainMiddleDisplay = "main middle"
const CreateCalendarForm = "create calendar form"
const EditCalendarForm = "edit calendar form"

func (g *GUI) mainScreenMiddleCalendarDisplay() fyne.CanvasObject {
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

func (g *GUI) generateMainScreenMiddleDisplay(display string) {
	switch display {
	case MainMiddleDisplay:
		g.MiddleContainer.Objects = []fyne.CanvasObject{
			g.mainScreenMiddleCalendarDisplay(),
		}

		g.MiddleContainer.Refresh()
	case CreateCalendarForm:
		g.MiddleContainer.Objects = []fyne.CanvasObject{
			g.showCreateCalendar(),
		}

		g.MiddleContainer.Refresh()
	case EditCalendarForm:
		err := g.checkCalendarSelected()
		if err != nil {
			g.showError("No calendar selected.", err)
			return
		}
		g.MiddleContainer.Objects = []fyne.CanvasObject{
			g.showEditCalendar(g.Calendar.ID),
		}

		g.MiddleContainer.Refresh()
	}

}
