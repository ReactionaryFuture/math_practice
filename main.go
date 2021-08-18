package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func main() {
	app := app.New()
	window := app.NewWindow("math_practice")

	data := binding.NewFloat()
	data.Set(30)
	entry := widget.NewEntryWithData(binding.FloatToStringWithFormat(data, "%v"))
	slide := widget.NewSliderWithData(1, 60, data)

	cont := container.NewBorder(nil, nil, nil, entry, slide)

	window.SetContent(cont)
	window.Resize(fyne.NewSize(300, 1))
	window.ShowAndRun()
}

func makeCheckboxCard(mathType string) (*widget.Card, *[]*widget.Check) {
	// this function makes a card for a math type (i.e. + - *) with check
	// boxes that allow the inclusion or exclusion of various tables (i.e.
	// multiplication tables)

	// make check boxes for number sets 1 through 12 and set them to checked
	var chkBxs []*widget.Check
	for i := 0; i <= 12; i++ {
		temp := widget.NewCheck(strconv.Itoa(i)+"s", func(bool) {})
		temp.SetChecked(true)
		chkBxs = append(chkBxs, temp)
	}

	// make button to check all boxes
	checkAllBoxesBtn := widget.NewButton("Check all", func() {
		for _, chkBx := range chkBxs {
			chkBx.SetChecked(true)
		}
	})

	// make button to uncheck all boxes
	uncheckAllBoxesBtn := widget.NewButton("Uncheck all", func() {
		for _, chkBx := range chkBxs {
			chkBx.SetChecked(false)
		}
	})

	// put both buttons in a container with horizontal layout
	btnCont := container.NewGridWithColumns(
		2, checkAllBoxesBtn, uncheckAllBoxesBtn)

	// put all check boxes in a container in a grid pattern
	vBox1 := container.NewGridWithColumns(
		4, chkBxs[1], chkBxs[2], chkBxs[3], chkBxs[4])
	vBox2 := container.NewGridWithColumns(
		4, chkBxs[5], chkBxs[6], chkBxs[7], chkBxs[8])
	vBox3 := container.NewGridWithColumns(
		4, chkBxs[9], chkBxs[10], chkBxs[11], chkBxs[12])
	checkBoxCont := container.NewGridWithRows(3, vBox1, vBox2, vBox3)

	// put the button container and the check box container together in a
	// vertical layout
	finalLayout := container.NewVBox(btnCont, checkBoxCont)

	// wrap finalLayout in a card
	card := widget.NewCard(mathType, "", finalLayout)

	// return the card and the state map
	return card, &chkBxs
}

func makeSliderWithEntry()
