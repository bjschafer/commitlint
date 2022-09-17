package lint

import (
	"fmt"
	"strings"

	"github.com/bjschafer/commitlint/pkg/git"
)

func ExampleCondition() {
	condition := Condition{
		Part: git.CommitMessage{},
		Matcher: func(s string) bool {
			if strings.HasSuffix(s, ".") {
				return true
			}
			return false
		},
	}

	fmt.Printf("Condition matcher returned %v", condition.Matcher("hi."))
	// Output: Condition matcher returned true
}
