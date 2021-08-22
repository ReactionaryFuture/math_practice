package data

import "strconv"

type MathProblem struct {
	MathType          string
	Operator          string
	Ints              [2]int
	Answer            int
	AsString          string
	IsNegative        bool
	IsEasySubtraction bool
}

func MakeFullProblemsSet() [3][13][13]MathProblem {
	var mathProblems [3][13][13]MathProblem

	for i := 0; i <= 12; i++ {
		for j := 0; j <= 12; j++ {
			m := &mathProblems[0][i][j]
			m.MathType = "Addition"
			m.Operator = "+"
			m.Ints[0] = i
			m.Ints[1] = j
			m.Answer = m.Ints[0] + m.Ints[1]
			m.AsString = strconv.Itoa(i) + m.Operator + strconv.Itoa(j) + "="
			m.IsNegative = false
			m.IsEasySubtraction = false
		}
	}

	for i := 0; i <= 12; i++ {
		for j := 0; j <= 12; j++ {
			m := &mathProblems[1][i][j]
			m.MathType = "Subtraction"
			m.Operator = "-"
			m.Ints[0] = i
			m.Ints[1] = j
			m.Answer = m.Ints[0] - m.Ints[1]
			m.AsString = strconv.Itoa(i) + m.Operator + strconv.Itoa(j) + "="
			if m.Answer < 0 {
				m.IsNegative = true
			} else {
				m.IsNegative = false
			}
			if !m.IsNegative && m.Answer < 2 {
				m.IsEasySubtraction = true
			} else {
				m.IsEasySubtraction = false
			}
		}
	}

	for i := 0; i <= 12; i++ {
		for j := 0; j <= 12; j++ {
			m := &mathProblems[2][i][j]
			m.MathType = "Multiplication"
			m.Operator = "x"
			m.Ints[0] = i
			m.Ints[1] = j
			m.Answer = m.Ints[0] * m.Ints[1]
			m.AsString = strconv.Itoa(i) + m.Operator + strconv.Itoa(j) + "="
			m.IsNegative = false
			m.IsEasySubtraction = false
		}
	}

	return mathProblems
}
