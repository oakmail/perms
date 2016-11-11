package whitespace_test

import (
	"testing"

	"github.com/stratexio/perms/whitespace"
)

func TestWhitespace(t *testing.T) {
	str := " train "
	if !whitespace.ContainsWhitespace(str) {
		t.Fatalf("%q should contain whitespace", str)
	}

	str = "train"
	if whitespace.ContainsWhitespace(str) {
		t.Fatalf("%q should not contain whitespace", str)
	}

	str = "   \n\n"

	if !whitespace.OnlyWhitespace(str) {
		t.Fatalf("%q should be only whitespace", str)
	}
}
