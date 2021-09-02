package ui

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

type OptionsScreen struct {
	SelectedForInclusion CheckBoxes
	NumOfProbs           SliderAndEntry
	TimeLimit            SliderAndEntry
	Penalty              SliderAndEntry
	Start                *widget.Button
	CanvasObject         *fyne.Container
}

type CheckBoxes struct {
	AddChkBxs []*widget.Check
	SubChkBxs []*widget.Check
	MltChkBxs []*widget.Check
}

type SliderAndEntry struct {
	Selection    binding.Float
	canvasObject *widget.Card
}

type ToggleIncludedGroup struct {
	checkBoxes   []*widget.Check
	canvasObject *widget.Card
}

func NewOptionsScreen() OptionsScreen {
	// this function makes an options screen that alows the user to select which math problems they want, how many
	// questions they want, how much time they want to solve each problem and how many times they want failed problems
	// to be reshuffled back in

	// make check and button groups
	additionToggles := newToggleIncludedGroup("Addition")
	subtractionToggles := newToggleIncludedGroup("Subtraction")
	multiplicationToggles := newToggleIncludedGroup("Multiplication")

	// make slider and entry groups
	numOfProbsSlide := newSliderAndEntry(10, 90, 999, "How many problems?")
	timeLimitSlide := newSliderAndEntry(1, 19, 999, "Seconds per problem?")
	penaltySlide := newSliderAndEntry(1, 5, 99, "Reshuffle times?")

	//make start button
	startButton := widget.NewButton("Start", func() {})
	startButton.Importance = widget.HighImportance

	// make the canvas object
	canvasObject := container.NewBorder(
		nil,
		startButton,
		nil,
		nil,
		container.NewBorder(
			nil,
			nil,
			container.NewGridWithColumns(
				3,
				additionToggles.canvasObject,
				subtractionToggles.canvasObject,
				multiplicationToggles.canvasObject,
			),
			nil,
			container.NewGridWithRows(
				3,
				numOfProbsSlide.canvasObject,
				timeLimitSlide.canvasObject,
				penaltySlide.canvasObject,
			),
		),
	)

	// use border to get the sliders to be the expanding part
	checkBoxes := CheckBoxes{
		AddChkBxs: additionToggles.checkBoxes,
		SubChkBxs: subtractionToggles.checkBoxes,
		MltChkBxs: multiplicationToggles.checkBoxes,
	}

	// make and return the OptionsScreen struct
	optionsScreen := OptionsScreen{
		SelectedForInclusion: checkBoxes,
		NumOfProbs:           numOfProbsSlide,
		TimeLimit:            timeLimitSlide,
		Penalty:              penaltySlide,
		Start:                startButton,
		CanvasObject:         canvasObject,
	}
	return optionsScreen
}

func newToggleIncludedGroup(mathType string) ToggleIncludedGroup {
	// this function makes a widget.Card for a math type (i.e. + - *) with check boxes that allow the inclusion or
	// exclusion of various tables (i.e. multiplication tables)

	// make check boxes for number sets 1 through 12 and set them all to checked
	var checkBoxes []*widget.Check
	for i := 0; i <= 11; i++ {
		checkBox := widget.NewCheck(strconv.Itoa(i+1)+"s", func(bool) {})
		checkBox.SetChecked(true)
		checkBoxes = append(checkBoxes, checkBox)
	}

	// make buttons to check and uncheck all check boxes
	checkAllBoxesButton := widget.NewButton("Check all", func() {
		for _, checkBox := range checkBoxes {
			checkBox.SetChecked(true)
		}
	})
	uncheckAllBoxesButton := widget.NewButton("Uncheck all", func() {
		for _, checkBox := range checkBoxes {
			checkBox.SetChecked(false)
		}
	})

	// make the canvas object
	canvasObject := widget.NewCard(
		mathType,
		"",
		container.NewBorder(
			container.NewGridWithColumns(
				2,
				checkAllBoxesButton,
				uncheckAllBoxesButton,
			),
			nil,
			nil,
			nil,
			container.NewAdaptiveGrid(
				4,
				checkBoxes[0], checkBoxes[1], checkBoxes[2], checkBoxes[3],
				checkBoxes[4], checkBoxes[5], checkBoxes[6], checkBoxes[7],
				checkBoxes[8], checkBoxes[9], checkBoxes[10], checkBoxes[11],
			),
		),
	)

	// make and return the checkCard struct
	ToggleIncludedGroup := &ToggleIncludedGroup{
		checkBoxes:   checkBoxes,
		canvasObject: canvasObject,
	}
	return *ToggleIncludedGroup
}

func newSliderAndEntry(sliderMin, sliderMax, maxVal int, label string) SliderAndEntry {
	// this function returns a widget.Card with slider and entry widgets that are bound together

	// make slider and slider widgets that are bound toether
	sliderValue := binding.NewFloat()
	sliderValue.Set(float64((sliderMin + sliderMax) / 2))
	entry := widget.NewEntryWithData(
		binding.FloatToStringWithFormat(sliderValue, "%v"),
	)
	slider := widget.NewSliderWithData(
		float64(sliderMin),
		float64(sliderMax),
		sliderValue,
	)

	// check entry for soundness. set to half max for entries that can not be easily converted to ints and set to
	// maxValue for entries greater than the max value
	entry.OnChanged = func(s string) {
		i, err := strconv.Atoi(s)
		if s == "" {
			return
		} else if err != nil {
			sliderValue.Set(float64((sliderMin + sliderMax) / 2))
		} else if i > maxVal {
			sliderValue.Set(float64(maxVal))
		} else {
			sliderValue.Set(float64(i))
		}
	}

	// combine the slider and entry boxes in a horizontal layout and wrap them in a card widget
	canvasObject := widget.NewCard(
		"",
		label,
		container.NewBorder(
			nil,
			nil,
			nil,
			entry,
			slider,
		),
	)

	// make and return the sliderCard struct
	sliderAndEntryGroup := &SliderAndEntry{
		Selection:    sliderValue,
		canvasObject: canvasObject,
	}
	return *sliderAndEntryGroup
}
