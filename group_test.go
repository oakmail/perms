package perms

import "testing"

func TestNewGroup(t *testing.T) {
	gr := NewGroup("billy")

	if gr.Name != "billy" {
		t.Fatalf("gr.Name is %v", gr.Name)
	}

	if gr.Nodes == nil {
		t.Fatalf("gr.Nodes should be instantiated")
	}

	if gr.Parents == nil {
		t.Fatalf("gr.Parents should be instantiated")
	}

}
