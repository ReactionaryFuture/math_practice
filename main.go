package main

import (
	"fmt"
	"image/color"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// test

func main() {
	app := app.New()
	window := app.NewWindow("math_practice")
	optionsScreen := NewOptionsScreen()

	optionsScreen.startButton.OnTapped = func() {
		fmt.Println(makeProblems(optionsScreen.additionToggles))
	}

	window.SetContent(optionsScreen.canvasObject)
	window.ShowAndRun()
}

func makeProblems(toggles ToggleIncludedGroup) []string {
	var operator string
	switch toggles.canvasObject.Title {
	case "Addition":
		operator = "+"
	case "Subtraxtion":
		operator = "-"
	case "Multiplication":
		operator = "x"
	}

	var allChecked []string
	for _, e := range toggles.checkBoxes {
		if e.Checked {
			allChecked = append(allChecked, strings.Trim(e.Text, "s"))
		}
	}

	var problemsList []string
	for _, e1 := range allChecked {
		for _, e2 := range allChecked {
			problemsList = append(problemsList, (e1 + operator + e2))
		}
	}

	return problemsList
}

type OptionsScreen struct {
	additionToggles       ToggleIncludedGroup
	subtractionToggles    ToggleIncludedGroup
	multiplicationToggles ToggleIncludedGroup
	numberOfProblems      SliderAndEntryGroup
	timeLimitPerProblem   SliderAndEntryGroup
	ReshufflePenalty      SliderAndEntryGroup
	startButton           *widget.Button
	canvasObject          *fyne.Container
}

func NewOptionsScreen() OptionsScreen {
	// this function makes an options screen that alows the user to select which math problems they want, how many
	// questions they want, how much time they want to solve each problem and how many times they want failed problems
	// to be reshuffled back in

	// make check and button groups
	additionToggles := NewToggleIncludedGroup("Addition")
	subtractionToggles := NewToggleIncludedGroup("Subtraction")
	multiplicationToggles := NewToggleIncludedGroup("Multiplication")

	// make slider and entry groups
	numberOfProblems := NewSliderAndEntryGroup(10, 90, 999, "How many problems?")
	timeLimitPerProblem := NewSliderAndEntryGroup(1, 20, 999, "Seconds per problem?")
	ReshufflePenalty := NewSliderAndEntryGroup(1, 5, 99, "Reshuffle times?")

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
				&timeLimitPerProblem.canvasObject,
				&ReshufflePenalty.canvasObject,
			),
			startButton,
		),
		nil,
		nil,
		container.NewGridWithColumns(
			3,
			additionToggles.canvasObject,
			subtractionToggles.canvasObject,
			multiplicationToggles.canvasObject,
		),
	)

	// make and return the OptionsScreen struct
	optionsScreen := &OptionsScreen{
		additionToggles:       additionToggles,
		subtractionToggles:    subtractionToggles,
		multiplicationToggles: multiplicationToggles,
		numberOfProblems:      numberOfProblems,
		timeLimitPerProblem:   timeLimitPerProblem,
		ReshufflePenalty:      numberOfProblems,
		startButton:           startButton,
		canvasObject:          canvasObject,
	}
	return *optionsScreen
}

type QuestionScreen struct {
	timer        *TimerBar
	entry        *widget.Entry
	canvasObject *fyne.Container
}

func NewQuestionScreen(question string) QuestionScreen {
	// this function creates the interface where questions are administered and answers are given by the user

	// create a timer bar at the top of the page to show the user how much more time he has before the question is
	// marked incorrect.
	timer := NewTimerBar()
	timer.TextFormatter = func() string {
		return ""
	}

	//create the entry object where the user recieves his question and gives his answers
	entry := widget.NewEntry()
	entry.SetText(question)

	// make the canvas object
	canvasObject := container.NewBorder(timer, nil, nil, nil, entry)

	// make and return the QuestionScreen struct
	QuestionScreen := &QuestionScreen{
		timer:        timer,
		entry:        entry,
		canvasObject: canvasObject,
	}
	return *QuestionScreen
}

type ToggleIncludedGroup struct {
	checkBoxes   []*widget.Check
	canvasObject *widget.Card
}

func NewToggleIncludedGroup(mathType string) ToggleIncludedGroup {
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

type TimerBar struct {
	widget.ProgressBar
}

func NewTimerBar() *TimerBar {
	// timer bar is created as distinct from widget.ProgressBar because the default MinSize needs to be overridden
	// as a temporary hack to get large text in the entry widget until rich text is released.

	prog := &TimerBar{}
	prog.ExtendBaseWidget(prog)

	return prog
}

func (w *TimerBar) MinSize() fyne.Size {
	// override the minsize of widget.ProgressBar

	return fyne.NewSize(1, 30)
}

type BigTextTheme struct{}

func (m BigTextTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	return theme.DefaultTheme().Color(name, variant)
}

func (m BigTextTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

func (m BigTextTheme) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(style)
}

func (m BigTextTheme) Size(name fyne.ThemeSizeName) float32 {
	if name == theme.SizeNameText {
		return 200
	} else {
		return theme.DefaultTheme().Size(name)
	}
}
