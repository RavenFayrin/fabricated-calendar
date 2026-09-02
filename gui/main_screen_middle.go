package gui

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
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
	err = g.checkCalendarData()
	if err != nil {
		content := container.NewVBox(
			widget.NewLabelWithStyle(
				"No Months or Weekdays Created",
				fyne.TextAlignCenter,
				fyne.TextStyle{Bold: true},
			),
			widget.NewLabelWithStyle(
				"Create months and/or weekdays to begin.",
				fyne.TextAlignCenter,
				fyne.TextStyle{Bold: false},
			),
		)
		return content
	}

	backMonthButton := widget.NewButtonWithIcon(
		"",
		theme.NavigateBackIcon(),
		func() {
			err := g.previousMonth()
			if err != nil {
				g.showError("Unable to show previous month.", err)
			}
		})

	nextMonthButton := widget.NewButtonWithIcon(
		"",
		theme.NavigateNextIcon(),
		func() {
			err := g.nextMonth()
			if err != nil {
				g.showError("Unable to show next month.", err)
			}
		})

	weekdayGrid := g.createWeekdayGrid()
	monthGrid := g.createMonthGrid()

	content := container.NewPadded(
		container.NewVBox(
			widget.NewLabelWithStyle(
				g.Calendar.Name,
				fyne.TextAlignCenter,
				fyne.TextStyle{Bold: true},
			),
			container.NewHBox(
				layout.NewSpacer(),
				backMonthButton,
				widget.NewLabelWithStyle(
					fmt.Sprintf("%s - Year %v", g.CalendarData.Months[g.DisplayMonthIndex].Name, g.DisplayYear),
					fyne.TextAlignCenter,
					fyne.TextStyle{Bold: true},
				),
				nextMonthButton,
				layout.NewSpacer(),
			),
			weekdayGrid,
			monthGrid,
		))

	return content
}

func (g *GUI) nextMonth() error {
	err := g.checkCalendarSelected()
	if err != nil {
		return err
	}

	g.DisplayMonthIndex++

	if g.DisplayMonthIndex >= int32(len(g.CalendarData.Months)) {
		g.DisplayMonthIndex = 0
		g.DisplayYear++
	}

	g.generateMainScreenMiddleDisplay(MainMiddleDisplay)

	return nil
}

func (g *GUI) previousMonth() error {
	err := g.checkCalendarSelected()
	if err != nil {
		return err
	}

	g.DisplayMonthIndex--

	if g.DisplayMonthIndex < 0 {
		if g.DisplayYear == 0 {
			g.DisplayMonthIndex = 0
		}

		g.DisplayYear--
		g.DisplayMonthIndex = int32(len(g.CalendarData.Months)) - 1
	}

	g.generateMainScreenMiddleDisplay(MainMiddleDisplay)

	return nil
}

func (g *GUI) createWeekdayGrid() fyne.CanvasObject {
	weekdayGrid := container.New(
		layout.NewGridLayout(len(g.CalendarData.Weekdays)))

	for _, weekday := range g.CalendarData.Weekdays {
		weekdayLable := widget.NewLabelWithStyle(
			weekday.Name,
			fyne.TextAlignCenter,
			fyne.TextStyle{Bold: true},
		)

		weekdayGrid.Add(weekdayLable)
	}
	return weekdayGrid
}

func (g *GUI) createMonthGrid() fyne.CanvasObject {
	columns := len(g.CalendarData.Weekdays)

	monthGrid := container.New(
		layout.NewGridLayout(columns),
	)

	startWeekdayIndex := g.CalendarData.MonthStartWeekday(
		int(g.DisplayMonthIndex),
		int(g.DisplayYear),
	)

	// Blank cells before the first day.
	for range startWeekdayIndex {
		monthGrid.Add(NewCalendarCell(""))
	}

	// Days in the month.
	for i := range g.CalendarData.Months[g.DisplayMonthIndex].DaysInMonth {
		day := strconv.Itoa(int(i) + 1)

		monthGrid.Add(
			NewCalendarCell(day),
		)
	}

	// Blank cells after the last day so the final row is complete.
	totalCells := startWeekdayIndex +
		int(g.CalendarData.Months[g.DisplayMonthIndex].DaysInMonth)

	remainder := totalCells % columns

	if remainder != 0 {
		blanksNeeded := columns - remainder

		for range blanksNeeded {
			monthGrid.Add(NewCalendarCell(""))
		}
	}

	return monthGrid
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
