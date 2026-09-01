package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type CalendarCell struct {
	widget.BaseWidget
	Day string
}

func NewCalendarCell(day string) *CalendarCell {
	cell := &CalendarCell{
		Day: day,
	}

	cell.ExtendBaseWidget(cell)

	return cell
}

func (c *CalendarCell) CreateRenderer() fyne.WidgetRenderer {
	background := canvas.NewRectangle(
		theme.Color(theme.ColorNameInputBackground),
	)

	dayLabel := widget.NewLabelWithStyle(
		c.Day,
		fyne.TextAlignLeading,
		fyne.TextStyle{
			Bold: false,
		},
	)

	content := container.NewPadded(dayLabel)

	objects := []fyne.CanvasObject{
		background,
		content,
	}

	return &calendarCellRenderer{
		cell:       c,
		background: background,
		dayLabel:   dayLabel,
		objects:    objects,
	}
}

type calendarCellRenderer struct {
	cell       *CalendarCell
	background *canvas.Rectangle
	dayLabel   *widget.Label
	objects    []fyne.CanvasObject
}

func (r *calendarCellRenderer) Layout(size fyne.Size) {
	r.background.Resize(size)

	padding := theme.Padding()

	r.dayLabel.Move(fyne.NewPos(
		padding,
		padding,
	))

	r.dayLabel.Resize(fyne.NewSize(
		size.Width-(padding*2),
		size.Height-(padding*2),
	))
}

func (r *calendarCellRenderer) MinSize() fyne.Size {
	return fyne.NewSize(
		100,
		100,
	)
}

func (r *calendarCellRenderer) Objects() []fyne.CanvasObject {
	return r.objects
}

func (r *calendarCellRenderer) Destroy() {}

func (r *calendarCellRenderer) Refresh() {
	r.background.FillColor = theme.Color(
		theme.ColorNameInputBackground,
	)

	r.dayLabel.Text = r.cell.Day

	r.background.Refresh()
	r.dayLabel.Refresh()
}
