package registry

import "github.com/bjschafer/commitlint/pkg/lint"

type Registry interface {
	Register(*lint.Problem)
	All() []lint.Problem
}

type registry struct {
	lints map[string]lint.Problem
}

func NewRegistry() Registry {
	return registry{}
}

func (r registry) Register(problem *lint.Problem) {
	r.lints[problem.Name] = *problem
}

func (r registry) All() []lint.Problem {
	ret := make([]lint.Problem, len(r.lints))

	i := 0
	for _, l := range r.lints {
		ret[i] = l
		i++
	}

	return ret
}
