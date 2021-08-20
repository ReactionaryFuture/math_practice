package main

import (
	"image/color"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	app := app.New()
	window := app.NewWindow("math_practice")
	app.Settings().SetTheme(&myTheme{})

	test := NewQuestionScreen()
	window.SetContent(test.canvasObject)

	window.Resize(fyne.NewSize(800, 300))
	window.SetContent(test.canvasObject)
	window.ShowAndRun()
}

type QuestionScreen struct {
	timerBar     widget.ProgressBar
	entry        widget.Entry
	canvasObject *fyne.Container
}

func NewQuestionScreen() QuestionScreen {
	timerBar := widget.NewProgressBar()
	entry := widget.NewEntry()
	timerBar.TextFormatter = func() string {
		return ""
	}

	canvasObject := container.NewBorder(timerBar, nil, nil, nil, entry)

	questionScreen := &QuestionScreen{
		timerBar:     *timerBar,
		entry:        *entry,
		canvasObject: canvasObject,
	}
	return *questionScreen
}

type OptionsScreen struct {
	additionToggles       CheckAndButtonGroup
	subtractionToggles    CheckAndButtonGroup
	multiplicationToggles CheckAndButtonGroup
	numberOfProblems      SliderAndEntryGroup
	timeLimitPerProblem   SliderAndEntryGroup
	penaltyReshuffle      SliderAndEntryGroup
	startButton           *widget.Button
	canvasObject          *fyne.Container
}

func NewOptionsScreen() OptionsScreen {
	// this function makes an options screen that alows the user to select which math problems they want, how many
	// questions they want, how much time they want to solve each problem and how many times they want failed problems
	// to be reshuffled back in

	// make check and button groups
	additionToggles := NewCheckAndButtonGroup("Addition")
	subtractionToggles := NewCheckAndButtonGroup("Subtraction")
	multiplicationToggles := NewCheckAndButtonGroup("Multiplication")

	// make slider and entry groups
	numberOfProblems := NewSliderAndEntryGroup(10, 90, 999, "How many problems?")
	timeLimitPerProblem := NewSliderAndEntryGroup(1, 20, 999, "Seconds per problem?")
	penaltyReshuffle := NewSliderAndEntryGroup(1, 5, 99, "Reshuffle times?")

	//make start button
	startButton := widget.NewButton("Start", func() {})
	startButton.Importance = widget.HighImportance

	// make the canvas object
	canvasObject := container.NewBorder(
		nil,
		container.NewVBox(
			container.NewGridWithColumns(
				3,
				&numberOfProblems.canvasObject,
				&penaltyReshuffle.canvasObject,
				&penaltyReshuffle.canvasObject,
			),
			startButton,
		),
		nil,
		nil,
		container.NewGridWithColumns(
			3,
			&additionToggles.canvasObject,
			&subtractionToggles.canvasObject,
			&multiplicationToggles.canvasObject,
		),
	)

	// make and return the OptionsScreen struct
	optionsScreen := &OptionsScreen{
		additionToggles:       additionToggles,
		subtractionToggles:    subtractionToggles,
		multiplicationToggles: multiplicationToggles,
		numberOfProblems:      numberOfProblems,
		timeLimitPerProblem:   timeLimitPerProblem,
		penaltyReshuffle:      numberOfProblems,
		startButton:           startButton,
		canvasObject:          canvasObject,
	}
	return *optionsScreen
}

type CheckAndButtonGroup struct {
	checkBoxes   []*widget.Check
	canvasObject widget.Card
}

func NewCheckAndButtonGroup(mathType string) CheckAndButtonGroup {
	// this function makes a widget.Card for a math type (i.e. + - *) with check boxes that allow the inclusion or
	// exclusion of various tables (i.e. multiplication tables)

	// make check boxes for number sets 1 through 12 and set them all to checked
	var checkBoxes []*widget.Check
	for i := 0; i <= 12; i++ {
		temp := widget.NewCheck(strconv.Itoa(i)+"s", func(bool) {})
		temp.SetChecked(true)
		checkBoxes = append(checkBoxes, temp)
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
				checkBoxes[1], checkBoxes[2], checkBoxes[3], checkBoxes[4],
				checkBoxes[5], checkBoxes[6], checkBoxes[7], checkBoxes[8],
				checkBoxes[9], checkBoxes[10], checkBoxes[11], checkBoxes[12],
			),
		),
	)

	// make and return the checkCard struct
	checkAndButtonGroup := &CheckAndButtonGroup{
		checkBoxes:   checkBoxes,
		canvasObject: *canvasObject,
	}
	return *checkAndButtonGroup
}

type SliderAndEntryGroup struct {
	selection    binding.Float
	canvasObject widget.Card
}

func NewSliderAndEntryGroup(sliderMin, sliderMax, maxVal int, label string) SliderAndEntryGroup {
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
	sliderAndEntryGroup := &SliderAndEntryGroup{
		selection:    sliderValue,
		canvasObject: *canvasObject,
	}
	return *sliderAndEntryGroup
}

type NoTextProgressBar struct {
	widget.BaseWidget

	Min, Max, Value float64

	valueSource   binding.Float
	valueListener binding.DataListener
}

type myTheme struct{}

func (m myTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	return theme.DefaultTheme().Color(name, variant)
}

func (m myTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

func (m myTheme) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(style)
}

func (m myTheme) Size(name fyne.ThemeSizeName) float32 {
	if name == theme.SizeNameText {
		return 60
	} else {
		return theme.DefaultTheme().Size(name)
	}
}
