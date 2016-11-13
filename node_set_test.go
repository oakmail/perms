package perms

import "testing"

func TestCheckNode(t *testing.T) {
	u := NewUser("ammar")

	u.Nodes = []Node{MustParseNode("projects.test"), MustParseNode("projects.build"), MustParseNode("projects.chat")}

	matched, negated := Check(u, MustParseNode("projects.test"))

	if negated {
		t.Fatalf("negated should be false")
	}

	if !matched {
		t.Fatalf("matched should be true")
	}

	u.Nodes = append(u.Nodes, MustParseNode("-projects.test"))

	matched, negated = Check(u, MustParseNode("projects.test"))

	if !negated {
		t.Fatalf("negated should be true")
	}

	if !matched {
		t.Fatalf("matched should be true")
	}

}
