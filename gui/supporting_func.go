package gui

import (
	"errors"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
)

func (g *GUI) showError(userMessage string, err error) {
	fyne.LogError(userMessage, err)
	dialog.ShowError(errors.New(userMessage), g.Window)
}
