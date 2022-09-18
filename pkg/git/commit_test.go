package git

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestParseFromString(t *testing.T) {
	tests := []struct {
		name        string
		giveMessage string
		wantMessage CommitMessage
		wantError   error
	}{
		{
			name:        "subject only",
			giveMessage: "hello there world i'm a subject :)",
			wantMessage: CommitMessage{Subject: "hello there world i'm a subject :)"},
		},
		{
			name: "subject, body, and single-line footer",
			giveMessage: `hello there i'm a subject

this is part of
a delicious breakfast bod

Bug: DADBOD-1234`,
			wantMessage: CommitMessage{
				Subject: "hello there i'm a subject",
				Body:    []string{"this is part of", "a delicious breakfast bod"},
				Footer:  []string{"Bug: DADBOD-1234"},
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		actualMsg, actualError := ParseFromString(tt.giveMessage)
		if diff := cmp.Diff(tt.wantMessage, actualMsg); diff != "" {
			t.Errorf("message didn't match. (-got +want)\n%s", diff)
		}
		if diff := cmp.Diff(tt.wantError, actualError); diff != "" {
			t.Errorf("error didn't match. (-got +want)\n%s", diff)
		}
	}
}

func TestString(t *testing.T) {
	tests := []struct {
		name        string
		giveMessage string
		wantError   error
	}{
		{
			name: "subject, body, single-line footer",
			giveMessage: `hello there world i'm a subject

i am a commit
that does some stuff

signed-off-by: me`,
		},
	}

	for _, tt := range tests {
		tt := tt

		cm, err := ParseFromString(tt.giveMessage)
		if diff := cmp.Diff(tt.giveMessage, cm.String()); diff != "" {
			t.Errorf("stringing didn't match. (-got +want)\n%s", diff)
		}
		if diff := cmp.Diff(tt.wantError, err); diff != "" {
			t.Errorf("error didn't match. (-got +want)\n%s", diff)
		}
	}
}
