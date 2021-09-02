package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type QuestionScreen struct {
	Progress     *noTextProgBar
	Timer        *noTextProgBar
	Entry        *widget.Entry
	CanvasObject *fyne.Container
}

type noTextProgBar struct {
	widget.ProgressBar
}

func NewQuestionScreen(initialQuestion string) QuestionScreen {
	// this function creates the interface where questions are administered and answers are given by the user

	// create a bar at the top of the page to show the user how much progress he has on answering all of the questions.
	progBar := newNoTextProgBar()
	progBar.TextFormatter = func() string {
		return ""
	}

	// create a progress bar at the top of the page to show the user how much more time he has before the question is
	// marked incorrect
	timerBar := newNoTextProgBar()
	timerBar.TextFormatter = func() string {
		return ""
	}

	//create the entry object where the user recieves his question and gives his answers
	entry := widget.NewEntry()
	entry.SetText(initialQuestion)

	// make the canvas object
	bars := container.NewVBox(progBar, timerBar)
	canvasObject := container.NewBorder(bars, nil, nil, nil, entry)

	// make and return the QuestionScreen struct
	QuestionScreen := &QuestionScreen{
		Progress:     progBar,
		Timer:        timerBar,
		Entry:        entry,
		CanvasObject: canvasObject,
	}
	return *QuestionScreen
}

func newNoTextProgBar() *noTextProgBar {
	// timer bar is created as distinct from widget.ProgressBar because the default MinSize needs to be overridden
	// as a temporary hack to get large text in the entry widget until rich text is released.

	prog := &noTextProgBar{}
	prog.ExtendBaseWidget(prog)

	return prog
}

func (w *noTextProgBar) MinSize() fyne.Size {
	// override the minsize of widget.ProgressBar

	return fyne.NewSize(1, 30)
}
