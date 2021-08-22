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

var allAdditionProblems = []string{
	"1+1=", "1+2=", "1+3=", "1+4=", "1+5=", "1+6=", "1+7=", "1+8=", "1+9=", "1+10=", "1+11=", "1+12=",
	"2+1=", "2+2=", "2+3=", "2+4=", "2+5=", "2+6=", "2+7=", "2+8=", "2+9=", "2+10=", "2+11=", "2+12=",
	"3+1=", "3+2=", "3+3=", "3+4=", "3+5=", "3+6=", "3+7=", "3+8=", "3+9=", "3+10=", "3+11=", "3+12=",
	"4+1=", "4+2=", "4+3=", "4+4=", "4+5=", "4+6=", "4+7=", "4+8=", "4+9=", "4+10=", "4+11=", "4+12=",
	"5+1=", "5+2=", "5+3=", "5+4=", "5+5=", "5+6=", "5+7=", "5+8=", "5+9=", "5+10=", "5+11=", "5+12=",
	"6+1=", "6+2=", "6+3=", "6+4=", "6+5=", "6+6=", "6+7=", "6+8=", "6+9=", "6+10=", "6+11=", "6+12=",
	"7+1=", "7+2=", "7+3=", "7+4=", "7+5=", "7+6=", "7+7=", "7+8=", "7+9=", "7+10=", "7+11=", "7+12=",
	"8+1=", "8+2=", "8+3=", "8+4=", "8+5=", "8+6=", "8+7=", "8+8=", "8+9=", "8+10=", "8+11=", "8+12=",
	"9+1=", "9+2=", "9+3=", "9+4=", "9+5=", "9+6=", "9+7=", "9+8=", "9+9=", "9+10=", "9+11=", "9+12=",
	"10+1=", "10+2=", "10+3=", "10+4=", "10+5=", "10+6=", "10+7=", "10+8=", "10+9=", "10+10=", "10+11=", "10+12=",
	"11+1=", "11+2=", "11+3=", "11+4=", "11+5=", "11+6=", "11+7=", "11+8=", "11+9=", "11+10=", "11+11=", "11+12=",
	"12+1=", "12+2=", "12+3=", "12+4=", "12+5=", "12+6=", "12+7=", "12+8=", "12+9=", "12+10=", "12+11=", "12+12=",
}

var allSubtractionProblems = []string{
	"1-1=", "1-2=", "1-3=", "1-4=", "1-5=", "1-6=", "1-7=", "1-8=", "1-9=", "1-10=", "1-11=", "1-12=",
	"2-1=", "2-2=", "2-3=", "2-4=", "2-5=", "2-6=", "2-7=", "2-8=", "2-9=", "2-10=", "2-11=", "2-12=",
	"3-1=", "3-2=", "3-3=", "3-4=", "3-5=", "3-6=", "3-7=", "3-8=", "3-9=", "3-10=", "3-11=", "3-12=",
	"4-1=", "4-2=", "4-3=", "4-4=", "4-5=", "4-6=", "4-7=", "4-8=", "4-9=", "4-10=", "4-11=", "4-12=",
	"5-1=", "5-2=", "5-3=", "5-4=", "5-5=", "5-6=", "5-7=", "5-8=", "5-9=", "5-10=", "5-11=", "5-12=",
	"6-1=", "6-2=", "6-3=", "6-4=", "6-5=", "6-6=", "6-7=", "6-8=", "6-9=", "6-10=", "6-11=", "6-12=",
	"7-1=", "7-2=", "7-3=", "7-4=", "7-5=", "7-6=", "7-7=", "7-8=", "7-9=", "7-10=", "7-11=", "7-12=",
	"8-1=", "8-2=", "8-3=", "8-4=", "8-5=", "8-6=", "8-7=", "8-8=", "8-9=", "8-10=", "8-11=", "8-12=",
	"9-1=", "9-2=", "9-3=", "9-4=", "9-5=", "9-6=", "9-7=", "9-8=", "9-9=", "9-10=", "9-11=", "9-12=",
	"10-1=", "10-2=", "10-3=", "10-4=", "10-5=", "10-6=", "10-7=", "10-8=", "10-9=", "10-10=", "10-11=", "10-12=",
	"11-1=", "11-2=", "11-3=", "11-4=", "11-5=", "11-6=", "11-7=", "11-8=", "11-9=", "11-10=", "11-11=", "11-12=",
	"12-1=", "12-2=", "12-3=", "12-4=", "12-5=", "12-6=", "12-7=", "12-8=", "12-9=", "12-10=", "12-11=", "12-12=",
}

var allMultiplicationProblems = []string{
	"1x1=", "1x2=", "1x3=", "1x4=", "1x5=", "1x6=", "1x7=", "1x8=", "1x9=", "1x10=", "1x11=", "1x12=",
	"2x1=", "2x2=", "2x3=", "2x4=", "2x5=", "2x6=", "2x7=", "2x8=", "2x9=", "2x10=", "2x11=", "2x12=",
	"3x1=", "3x2=", "3x3=", "3x4=", "3x5=", "3x6=", "3x7=", "3x8=", "3x9=", "3x10=", "3x11=", "3x12=",
	"4x1=", "4x2=", "4x3=", "4x4=", "4x5=", "4x6=", "4x7=", "4x8=", "4x9=", "4x10=", "4x11=", "4x12=",
	"5x1=", "5x2=", "5x3=", "5x4=", "5x5=", "5x6=", "5x7=", "5x8=", "5x9=", "5x10=", "5x11=", "5x12=",
	"6x1=", "6x2=", "6x3=", "6x4=", "6x5=", "6x6=", "6x7=", "6x8=", "6x9=", "6x10=", "6x11=", "6x12=",
	"7x1=", "7x2=", "7x3=", "7x4=", "7x5=", "7x6=", "7x7=", "7x8=", "7x9=", "7x10=", "7x11=", "7x12=",
	"8x1=", "8x2=", "8x3=", "8x4=", "8x5=", "8x6=", "8x7=", "8x8=", "8x9=", "8x10=", "8x11=", "8x12=",
	"9x1=", "9x2=", "9x3=", "9x4=", "9x5=", "9x6=", "9x7=", "9x8=", "9x9=", "9x10=", "9x11=", "9x12=",
	"10x1=", "10x2=", "10x3=", "10x4=", "10x5=", "10x6=", "10x7=", "10x8=", "10x9=", "10x10=", "10x11=", "10x12=",
	"11x1=", "11x2=", "11x3=", "11x4=", "11x5=", "11x6=", "11x7=", "11x8=", "11x9=", "11x10=", "11x11=", "11x12=",
	"12x1=", "12x2=", "12x3=", "12x4=", "12x5=", "12x6=", "12x7=", "12x8=", "12x9=", "12x10=", "12x11=", "12x12=",
}

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
		temp := widget.NewCheck(strconv.Itoa(i+1)+"s", func(bool) {})
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
