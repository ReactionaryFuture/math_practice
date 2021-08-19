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

	test := NewOptionsScreen()

	window.SetContent(test.canvasObject)
	window.ShowAndRun()
}

type OptionsScreen struct {
	addCheckCard CheckCard
	subCheckCard CheckCard
	mltCheckCard CheckCard
	numOfProbs   SliderCard
	time         SliderCard
	reshufTimes  SliderCard
	startBtn     *widget.Button
	canvasObject *fyne.Container
}

func NewOptionsScreen() OptionsScreen {
	addCrd := NewCheckCard("Addition")
	subCrd := NewCheckCard("Subtraction")
	mltCrd := NewCheckCard("Multiplication")

	numOfProbs := NewSliderCard(10, 90, 999, "How many problems?")
	time := NewSliderCard(1, 20, 999, "Seconds per problem?")
	reshuf := NewSliderCard(1, 5, 99, "Reshuffle times?")

	startBtn := widget.NewButton("Start", func() {})
	startBtn.Importance = widget.HighImportance

	checkCrds := container.NewGridWithColumns(
		3, &addCrd.canvasObject, &subCrd.canvasObject, &mltCrd.canvasObject)
	sliders := container.NewGridWithColumns(
		3, &numOfProbs.canvasObject, &time.canvasObject, &reshuf.canvasObject)
	slidersAndStartBtn := container.NewVBox(sliders, startBtn)
	cont := container.NewBorder(nil, slidersAndStartBtn, nil, nil, checkCrds)

	optionsScreen := &OptionsScreen{
		addCheckCard: addCrd,
		subCheckCard: subCrd,
		mltCheckCard: mltCrd,
		numOfProbs:   numOfProbs,
		time:         time,
		reshufTimes:  reshuf,
		startBtn:     startBtn,
		canvasObject: cont,
	}
	return *optionsScreen
}

type CheckCard struct {
	checkBoxes   []*widget.Check
	canvasObject widget.Card
}

func NewCheckCard(mathType string) CheckCard {
	// this function makes a widget.Card for a math type (i.e. + - *)
	// with check boxes that allow the inclusion or exclusion of various
	// tables (i.e. multiplication tables)

	// make check boxes for number sets 1 through 12, set them to checked
	// and put all check boxes in a grid pattern container
	var chkBxs []*widget.Check
	for i := 0; i <= 12; i++ {
		temp := widget.NewCheck(strconv.Itoa(i)+"s", func(bool) {})
		temp.SetChecked(true)
		chkBxs = append(chkBxs, temp)
	}
	chkBxGrid := container.NewAdaptiveGrid(
		4,
		chkBxs[1], chkBxs[2], chkBxs[3], chkBxs[4],
		chkBxs[5], chkBxs[6], chkBxs[7], chkBxs[8],
		chkBxs[9], chkBxs[10], chkBxs[11], chkBxs[12],
	)

	// make button to check and uncheck all boxes and put the
	// the buttons in a horizontal layout
	chkAllBxsBtn := widget.NewButton("Check all", func() {
		for _, chkBx := range chkBxs {
			chkBx.SetChecked(true)
		}
	})
	unchkAllBxsBtn := widget.NewButton("Uncheck all", func() {
		for _, chkBx := range chkBxs {
			chkBx.SetChecked(false)
		}
	})
	btnCont := container.NewGridWithColumns(2, chkAllBxsBtn, unchkAllBxsBtn)

	// stack the buttons and check boxes in a vertical layout and wrap
	// them in a card widget
	combined := container.NewBorder(btnCont, nil, nil, nil, chkBxGrid)
	card := widget.NewCard(mathType, "", combined)

	// make and return the checkCard struct
	checkCard := &CheckCard{
		checkBoxes:   chkBxs,
		canvasObject: *card,
	}
	return *checkCard
}

type SliderCard struct {
	selection    binding.Float
	canvasObject widget.Card
}

func NewSliderCard(sliderMin, sliderMax, maxVal int, label string) SliderCard {
	// this function returns a widget.Card with slider and entry widgets
	// that are bound together

	// make slider and slider widgets that are bound toether
	value := binding.NewFloat()
	value.Set(float64((sliderMin + sliderMax) / 2))
	entry := widget.NewEntryWithData(
		binding.FloatToStringWithFormat(value, "%v"))
	slide := widget.NewSliderWithData(
		float64(sliderMin), float64(sliderMax), value)

	// check entry for soundness. set to half max for entries that can
	// not be easily converted to ints and set to maxValue for entries
	// greater than the max value
	entry.OnChanged = func(s string) {
		i, err := strconv.Atoi(s)
		if s == "" {
			return
		} else if err != nil {
			value.Set(float64((sliderMin + sliderMax) / 2))
		} else if i > maxVal {
			value.Set(float64(maxVal))
		} else {
			value.Set(float64(i))
		}
	}

	// combine the slider and entry boxes in a horizontal layout and wrap
	// them in a card widget
	combined := container.NewBorder(nil, nil, nil, entry, slide)
	card := widget.NewCard("", label, combined)

	// make and return the sliderCard struct
	sliderCard := &SliderCard{
		selection:    value,
		canvasObject: *card,
	}
	return *sliderCard
}
