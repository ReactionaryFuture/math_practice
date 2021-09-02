package data

import (
	"github/ReactionaryFuture/math_practice/base"
	"github/ReactionaryFuture/math_practice/generics"
	"github/ReactionaryFuture/math_practice/ui"
	"strconv"

	"fyne.io/fyne/v2/widget"
)

func NewQuestions(chkBxs ui.CheckBoxes, numOfQuestions int) []base.MathQuestion {
	additionQuestions :=
		newAddQuestions(
			newAllCombos(
				newListOfChecked(
					chkBxs.AddChkBxs,
				),
			),
		)

	subtractionQuestions :=
		newSubQuestion(
			newAllCombos(
				newListOfChecked(
					chkBxs.SubChkBxs,
				),
			),
		)

	multiplicationQuestions :=
		newMltQuestions(
			newAllCombos(
				newListOfChecked(
					chkBxs.MltChkBxs,
				),
			),
		)

	questions := additionQuestions
	questions = append(questions, subtractionQuestions...)
	questions = append(questions, multiplicationQuestions...)
	questions = shrink(questions, numOfQuestions)
	return questions
}

func newListOfChecked(cs []*widget.Check) []int {
	var acc []int
	for i, e := range cs {
		if e.Checked {
			acc = append(acc, i+1)
		}
	}
	return acc
}

func newAllCombos(is []int) []base.Params {
	var acc []base.Params
	for _, a := range is {
		for _, b := range is {
			acc = append(acc, base.Params{A: a, B: b})
		}
	}
	return acc
}

func newAddQuestions(ps []base.Params) []base.MathQuestion {
	qs := make([]base.MathQuestion, len(ps))
	for i, p := range ps {
		var (
			sQuestion = strconv.Itoa(p.A) + "+" + strconv.Itoa(p.B) + "="
			iAnswer   = p.A + p.B
			sAnswer   = sQuestion + strconv.Itoa(iAnswer)
		)
		qs[i] = base.MathQuestion{
			MathType:  "Addition",
			Operator:  "+",
			INums:     p,
			IAnswer:   iAnswer,
			SQuestion: sQuestion,
			SAnswer:   sAnswer,
			IsNeg:     false,
			IsTooEz:   false,
		}
	}
	return qs
}

func newSubQuestion(ps []base.Params) []base.MathQuestion {
	qs := make([]base.MathQuestion, len(ps))
	for i, p := range ps {
		var (
			sQuestion = strconv.Itoa(p.A) + "–" + strconv.Itoa(p.B) + "="
			iAnswer   = p.A - p.B
			sAnswer   = sQuestion + strconv.Itoa(iAnswer)
			isNeg     = p.A < 0
			isTooEz   = p.A-p.B == 1
		)
		qs[i] = base.MathQuestion{
			MathType:  "Subtraction",
			Operator:  "–",
			INums:     p,
			IAnswer:   iAnswer,
			SQuestion: sQuestion,
			SAnswer:   sAnswer,
			IsNeg:     isNeg,
			IsTooEz:   isTooEz,
		}
	}
	return qs
}

func newMltQuestions(ps []base.Params) []base.MathQuestion {
	qs := make([]base.MathQuestion, len(ps))
	for i, p := range ps {
		var (
			sQuestion = strconv.Itoa(p.A) + "×" + strconv.Itoa(p.B) + "="
			iAnswer   = p.A * p.B
			sAnswer   = sQuestion + strconv.Itoa(iAnswer)
		)
		qs[i] = base.MathQuestion{
			MathType:  "Mltiplication",
			Operator:  "×",
			INums:     p,
			IAnswer:   iAnswer,
			SQuestion: sQuestion,
			SAnswer:   sAnswer,
			IsNeg:     false,
			IsTooEz:   false,
		}
	}
	return qs
}

func shrink(qs []base.MathQuestion, n int) []base.MathQuestion {
	var acc []base.MathQuestion
	for i := 0; i < n; i++ {
		acc = append(acc, generics.RandQuestion(qs))
	}
	return acc
}
