package base

type Params struct{ A, B int }

type MathQuestion struct {
	MathType  string
	Operator  string
	INums     Params
	IAnswer   int
	SQuestion string
	SAnswer   string
	IsNeg     bool
	IsTooEz   bool
}
