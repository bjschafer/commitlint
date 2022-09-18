package lint

import "github.com/bjschafer/commitlint/pkg/git"

// Rule defines a particular lint
type Rule struct {
	Name string
	// ApplicableAlways will ensure Condition applies if true, and ensure Condition *doesn't* apply if false.
	ApplicableAlways bool
	Condition        Condition
}

// Condition is a check that's run against a particular part of a git.CommitMessage.
type Condition struct {
	Part git.MessagePart
	// Matcher is a function that accepts the commit message part and checks it. True means the condition matched.
	Matcher func(string) bool
}
