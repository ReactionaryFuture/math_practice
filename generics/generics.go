package generics

import (
	"github/ReactionaryFuture/math_practice/base"
	"math/rand"
)

func RandQuestion(qs []base.MathQuestion) base.MathQuestion {
	return qs[rand.Intn(len(qs))]
}

func RemMathQuestion(q base.MathQuestion, qs []base.MathQuestion) []base.MathQuestion {
	var acc []base.MathQuestion
	for i, e := range qs {
		if e == q {
			acc = append(acc, qs[i+1:]...)
			break
		}
		acc = append(acc, e)
	}
	return acc
}

func AppendNTimes(q base.MathQuestion, qs []base.MathQuestion, n int) []base.MathQuestion {
	for i := 0; i < n; i++ {
		qs = append(qs, q)
	}
	return qs
}
