package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type GratzScreen struct{ CanvasObject *fyne.Container }

func NewGratzScreen() GratzScreen {
	label := widget.NewLabel("Congratz!")
	canvasObject := container.NewCenter(label)

	return GratzScreen{canvasObject}
}
