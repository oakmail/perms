package perms

import (
	"os"
	"testing"
)

func TestWeb_PrettyDump(t *testing.T) {
	if !testing.Verbose() {
		return
	}

	web := NewWeb()

	web.AddGroup(NewGroup("admin"))
	web.AddGroup(NewGroup("moderator"))

	web.AddUser(&User{
		Name:   "ammar",
		Groups: []string{"admin", "moderator"},
		Nodes:  []Node{MustParseNode("testing.t")},
	})

	web.PrettyDump(os.Stdout)
}
