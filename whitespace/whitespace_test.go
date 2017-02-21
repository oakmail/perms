package whitespace_test

import (
	"testing"

	"github.com/trtlio/perms/whitespace"
)

func TestWhitespace(t *testing.T) {
	str := " train "
	if !whitespace.Contains(str) {
		t.Fatalf("%q should contain whitespace", str)
	}

	str = "train"
	if whitespace.Contains(str) {
		t.Fatalf("%q should not contain whitespace", str)
	}

	str = "   \n\n"

	if !whitespace.Only(str) {
		t.Fatalf("%q should be only whitespace", str)
	}

	str = whitespace.Strip(str)

	if str != "" {
		t.Fatalf("%q should be empty", str)
	}
}
