package perms

//User contains a user with permissions
type User struct {
	Name   string
	Groups []string
	Nodes  []Node
}

//NewUser returns a pointer to an instantiated user
func NewUser(name string) *User {
	return &User{
		Name:   name,
		Groups: make([]string, 0, 5),
		Nodes:  make([]Node, 0, 5),
	}
}

//GetNodes returns all users nodes
func (u *User) GetNodes() []Node {
	return u.Nodes
}
