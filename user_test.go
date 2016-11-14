package perms

import "testing"

func TestNewUser(t *testing.T) {
	u := NewUser("billy")

	if u.Name != "billy" {
		t.Fatalf("u.Name is %v", u.Name)
	}

	if u.Nodes == nil {
		t.Fatalf("u.Nodes should be instantiated")
	}

	if u.Groups == nil {
		t.Fatalf("u.Groups should be instantiated")
	}
}
