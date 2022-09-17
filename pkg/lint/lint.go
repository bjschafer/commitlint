package lint

type Linter interface {
	Run(message string) []Problem
}

type linter struct {
}

func NewLinter() Linter {
	return linter{}
}

func (l linter) Run(message string) []Problem {
	var ret []Problem

	return ret
}