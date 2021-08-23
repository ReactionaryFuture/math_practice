package main

import (
	"image/color"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

func main() {
	app := app.New()
	window := app.NewWindow("math_practice")

	window.SetContent(canvas.NewRectangle(color.White))
	window.ShowAndRun()
}
