package main

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

func main() {
	app := app.New()
	window := app.NewWindow("math_practice")

	temp := []int{3, 1, 8}
	fmt.Println(AllIntPairs(temp))

	window.SetContent(canvas.NewRectangle(color.White))
	window.ShowAndRun()
}
