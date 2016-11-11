package perms

//Group contains a permission group
type Group struct {
	Name    string
	Parents []*Group `json:"parents"`
	Nodes   []Node   `json:"nodes"`
}
