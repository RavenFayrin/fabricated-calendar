package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
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
