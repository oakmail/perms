package perms

//Group contains a permission group
type Group struct {
	Name    string
	Parents []string `json:"parents"`
	Nodes   Nodes    `json:"nodes"`
}

//NewGroup returns a pointer to an instantied group
func NewGroup(name string) *Group {
	return &Group{
		Name:    name,
		Parents: make([]string, 0, 5),
		Nodes:   make(Nodes, 0, 5),
	}
}
