package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	xwidget "fyne.io/x/fyne/widget"
)

type LabelOptions struct {
	Alignment     fyne.TextAlign
	Bold          bool
	Italic        bool
	Monospace     bool
	Underline     bool
	Strikethrough bool
}

func textLabel(text string, options ...LabelOptions) *widget.Label {
	opts := LabelOptions{
		Alignment: fyne.TextAlignLeading,
	}

	if len(options) > 0 {
		opts = options[0]
	}

	return widget.NewLabelWithStyle(
		text,
		opts.Alignment,
		fyne.TextStyle{
			Bold:          opts.Bold,
			Italic:        opts.Italic,
			Monospace:     opts.Monospace,
			Underline:     opts.Underline,
			Strikethrough: opts.Strikethrough,
		},
	)
}

type ButtonOptions struct {
	Text          string
	Icon          fyne.Resource
	Alignment     widget.ButtonAlign
	Disabled      bool
	IconAlignment widget.ButtonIconPlacement
}

func button(function func(), options ...ButtonOptions) *widget.Button {
	opts := ButtonOptions{
		Text: "",
	}

	if len(options) > 0 {
		opts = options[0]
	}

	var button *widget.Button

	if opts.Icon != nil {
		button = widget.NewButtonWithIcon(opts.Text, opts.Icon, function)
	} else {
		button = widget.NewButton(opts.Text, function)
	}

	button.Alignment = opts.Alignment
	button.IconPlacement = opts.IconAlignment

	if opts.Disabled {
		button.Disable()
	}

	return button
}

type EntryOptions struct {
	PlaceHolder string
	MultiLine   bool
	Password    bool
	Disabled    bool
	OnSubmit    func(string)
}

func entry(options ...EntryOptions) *widget.Entry {
	opts := EntryOptions{}

	if len(options) > 0 {
		opts = options[0]
	}

	var entry *widget.Entry

	if opts.MultiLine {
		entry = widget.NewMultiLineEntry()
	} else if opts.Password {
		entry = widget.NewPasswordEntry()
	} else {
		entry = widget.NewEntry()
	}

	entry.SetPlaceHolder(opts.PlaceHolder)

	if opts.Disabled {
		entry.Disable()
	}

	entry.OnSubmitted = opts.OnSubmit

	return entry
}

func entryNumerical(options ...EntryOptions) *xwidget.NumericalEntry {
	opts := EntryOptions{}

	if len(options) > 0 {
		opts = options[0]
	}

	entry := xwidget.NewNumericalEntry()

	entry.SetPlaceHolder(opts.PlaceHolder)

	if opts.Disabled {
		entry.Disable()
	}

	return entry
}
