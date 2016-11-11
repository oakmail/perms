package perms

//User contains a user with permissions
type User struct {
	Name   string
	Groups []*Group
	Nodes  []Node
}
