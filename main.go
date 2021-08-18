package main

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func main() {
	app := app.New()
	window := app.NewWindow("math_practice")

	addCrd := mkNewCheckCard("Addition")
	subCrd := mkNewCheckCard("Subtraction")
	mltCrd := mkNewCheckCard("Multiplication")

	numProbsSld := mkNewSliderCard(10, 90, 999, "How many problems?")
	timeSld := mkNewSliderCard(1, 20, 999, "Seconds per problem?")
	reShufSld := mkNewSliderCard(1, 5, 99, "Reshuffle times?")

	startBtn := widget.NewButton("Start", func() {})
	startBtn.Importance = widget.HighImportance

	checkCrds := container.NewGridWithColumns(3, addCrd, subCrd, mltCrd)
	sliders := container.NewGridWithColumns(3, numProbsSld, timeSld, reShufSld)
	slidersAndStartBtn := container.NewVBox(sliders, startBtn)
	cont := container.NewBorder(nil, slidersAndStartBtn, nil, nil, checkCrds)

	window.Resize(fyne.NewSize(870, 365))
	window.SetContent(cont)
	window.ShowAndRun()
}

func mkNewCheckCard(mathType string) *widget.Card {
	// this function makes a widget.Card for a math type (i.e. + - *)
	// with check boxes that allow the inclusion or exclusion of various
	// tables (i.e. multiplication tables)

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
	btnCont := container.NewGridWithColumns(2, checkAllBoxesBtn, uncheckAllBoxesBtn)

	// put all check boxes in a container in a grid pattern
	chkBxGrid := container.NewAdaptiveGrid(
		4,
		chkBxs[1], chkBxs[2], chkBxs[3], chkBxs[4],
		chkBxs[5], chkBxs[6], chkBxs[7], chkBxs[8],
		chkBxs[9], chkBxs[10], chkBxs[11], chkBxs[12],
	)

	// put the button container and the check box container together in a
	// vertical layout
	finalLayout := container.NewBorder(btnCont, nil, nil, nil, chkBxGrid)

	// wrap finalLayout in a card
	card := widget.NewCard(mathType, "", finalLayout)

	// return the card and the state map
	return card
}

func mkNewSliderCard(sliderMin, sliderMax, maxVal int, label string) *widget.Card {
	// this function returns a widget.Card with slider and entry widgets
	// that are bound together

	// make slider and slider widgets that are bound toether
	data := binding.NewFloat()
	data.Set(float64((sliderMin + sliderMax) / 2))
	entry := widget.NewEntryWithData(binding.FloatToStringWithFormat(data, "%v"))
	slide := widget.NewSliderWithData(float64(sliderMin), float64(sliderMax), data)

	// check entry for soundness. set to half max for entries that can
	// not be easily converted to ints and set to maxValue for entries
	// greater than the max value
	entry.OnChanged = func(s string) {
		i, err := strconv.Atoi(s)
		if s == "" {
			return
		} else if err != nil {
			data.Set(float64((sliderMin + sliderMax) / 2))
		} else if i > maxVal {
			data.Set(float64(maxVal))
		} else {
			data.Set(float64(i))
		}
	}

	// make a container with layout for the widgets
	cont := container.NewBorder(nil, nil, nil, entry, slide)

	// wrap container in card
	card := widget.NewCard("", label, cont)

	return card
}
