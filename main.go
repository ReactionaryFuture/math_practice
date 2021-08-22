package main

import (
	"github/ReactionaryFuture/math_practice/ui"

	"fyne.io/fyne/v2/app"
)

func main() {
	app := app.New()
	window := app.NewWindow("math_practice")

	optionsScreen := ui.NewOptionsScreen()

	window.SetContent(optionsScreen.CanvasObject)
	window.ShowAndRun()
}
