package perms

//User contains a user with permissions
type User struct {
	Name   string
	Groups []string
	Nodes  Nodes
}

//NewUser returns a pointer to an instantiated user
func NewUser(name string) *User {
	return &User{
		Name:   name,
		Groups: make([]string, 0, 5),
		Nodes:  make(Nodes, 0, 5),
	}
}
