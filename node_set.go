package perms

//NodeSet defines a kind of object that owns nodes
type NodeSet interface {
	GetNodes() []Node
}

//CheckNode checks a node set for a permission
func CheckNode(ns NodeSet, check Node) (matched bool, negated bool) {
	for _, node := range ns.GetNodes() {
		if node.Match(check) {
			matched = true
			if node.Negate {
				return true, true
			}
		}
	}
	return matched, false
}
