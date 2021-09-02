package main

import (
	"github/ReactionaryFuture/math_practice/data"
	gen "github/ReactionaryFuture/math_practice/generics"
	"github/ReactionaryFuture/math_practice/ui"
	"math/rand"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	app := app.New()
	win := app.NewWindow("math_practice")
	opt := ui.NewOptionsScreen()

	opt.Start.OnTapped = func() {
		rand.Seed(time.Now().UnixNano())

		numOfProbs, _ := opt.NumOfProbs.Selection.Get()
		timeLimit, _ := opt.TimeLimit.Selection.Get()
		penalty, _ := opt.Penalty.Selection.Get()

		// innitialize list of questions using the data from check boxes
		// and the "How many problems?" slider on the options screen.
		questions := data.NewQuestions(opt.SelectedForInclusion, int(numOfProbs))
		currentQuestion := gen.RandQuestion(questions)
		questionScreen := ui.NewQuestionScreen(currentQuestion.SQuestion)

		// changing the theme here is essentially part of a hack to get around fyne not having rich text yet
		app.Settings().SetTheme(&ui.BigTextTheme{})
		win.SetContent(questionScreen.CanvasObject)

		// these values are
		progBarNumerator := 0
		progBarDenominator := len(questions)

		timer := time.NewTimer(time.Second * time.Duration(timeLimit))
		go func() {
			for {
				<-timer.C
				questionScreen.Entry.OnSubmitted("")
			}
		}()
		ticker := time.NewTicker(time.Millisecond * 100)
		tickerCounter := 0
		go func() {
			for {
				<-ticker.C
				tickerCounter++
				questionScreen.Timer.SetValue(float64(tickerCounter) / float64(timeLimit*10))
			}
		}()

		// this block get exicuted when the user presses enter on the main entry widget on the questions screen
		questionScreen.Entry.OnSubmitted = func(answer string) {
			timer.Reset(time.Second * time.Duration(timeLimit))
			tickerCounter = 0

			// check the answer
			if answer == currentQuestion.SAnswer { // if correct
				questions = gen.RemMathQuestion(currentQuestion, questions)
				if len(questions) <= 0 {
					fyne.CurrentApp().Quit()
					time.Sleep(time.Second)
				}
				progBarNumerator++
			} else { // if incorrect
				questions = gen.AppendNTimes(currentQuestion, questions, int(penalty))
				progBarDenominator = progBarDenominator + int(penalty)
			}

			// update the progress bar
			questionScreen.Progress.SetValue(float64(progBarNumerator) / float64(progBarDenominator))

			// get new question and write it to screen
			currentQuestion = gen.RandQuestion(questions)
			questionScreen.Entry.SetText(currentQuestion.SQuestion)

		}
	}

	win.Resize(fyne.NewSize(1130, 365))
	win.SetContent(opt.CanvasObject)
	win.ShowAndRun()
}

// move slider bars to the right
// implement bar for progress
// implement bar for time
// implement slider inputs
// fix window size
// report correct/incorrect screen
// make functions into methods

// stretch goals
// animations and fun stuff
