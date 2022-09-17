package git

import (
	"strings"
)

type CommitMessage struct {
	// Subject is the first line of the commit message, and usually has strict formatting rules
	Subject string
	// Body is the majority of the commit message
	Body []string
	// Footer contains things like Jira BUG: FOO-1234 links as well as sign-offs.
	Footer []string
}

func ParseFromString(rawMessage string) (CommitMessage, error) {
	ret := CommitMessage{}
	lines := strings.Split(rawMessage, `\n`)
	numLines := len(lines)

	ret.Subject = lines[0]
	if numLines <= 2 {
		return ret, nil
	}

	// for now, footer is limited
	var temp []string
	for i := numLines - 1; i >= 0; i-- {
		line := lines[i]
		if strings.TrimSpace(line) == "" {
			break
		}
		// if signed-off-by or Bug: (?)
		temp = append(temp, line)
	}
	// reverse, reverse
	ret.Footer = make([]string, len(temp))
	for i, line := range temp {
		ret.Footer[len(temp)-1-i] = line
	}

	// by rule, blank line separates subject from body
	if strings.TrimSpace(lines[1]) == "" {
		ret.Body = lines[2:]
	}

	return ret, nil
}
