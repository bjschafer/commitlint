package git

import (
	"fmt"
	"regexp"
	"strings"
)

type MessagePart int

const (
	Subject MessagePart = iota
	Body
	Footer
)

// msgParser parses a git commit message into subject, body, and footer
var msgParser = regexp.MustCompile(`^(?P<Subject>(?m)(?:^.+$)+)(?-m)(?:(?:(?:\r?\n){2})?(?P<Body>(?m)(?:^.+$(?:\r?\n)?)+)?(?-m)(?:(?:\r?\n){2})?(?P<Footer>(?m)(?:^.+$(?:\r?\n)?)+))?(?-m)$`)

type CommitMessage struct {
	// Subject is the first line of the commit message, and usually has strict formatting rules
	Subject string
	// Body is the majority of the commit message
	Body []string
	// Footer contains things like Jira BUG: FOO-1234 links as well as sign-offs.
	Footer []string
}

func ParseFromString(rawMessage string) (CommitMessage, error) {
	result := make(map[string]string)
	match := msgParser.FindStringSubmatch(rawMessage)
	if match == nil {
		return CommitMessage{Body: []string{rawMessage}}, fmt.Errorf("poorly formatted commit message")
	}

	for i, name := range msgParser.SubexpNames() {
		if i != 0 && name != "" {
			result[name] = match[i]
		}
	}

	var body []string = nil
	if result["Body"] != "" {
		body = strings.Split(result["Body"], "\n")
	}
	var footer []string = nil
	if result["Footer"] != "" {
		footer = strings.Split(result["Footer"], "\n")
	}

	return CommitMessage{
		Subject: result["Subject"],
		Body:    body,
		Footer:  footer,
	}, nil
}

func (c CommitMessage) String() string {
	return fmt.Sprintf("%s\n\n%s\n\n%s", c.Subject, strings.Join(c.Body, "\n"), strings.Join(c.Footer, "n"))
}
