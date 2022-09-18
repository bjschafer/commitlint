package rules

import (
	"strings"

	"github.com/bjschafer/commitlint/pkg/git"
	"github.com/bjschafer/commitlint/pkg/lint"
)

var BodyFullStop = lint.Rule{
	Name:             "BodyFullStop",
	ApplicableAlways: false,
	Condition: lint.Condition{
		Part: git.Body,
		Matcher: func(s string) bool {
			return strings.HasSuffix(s, ".")
		},
	},
}
