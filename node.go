package permissions

import (
	"strings"

	"github.com/pkg/errors"
	"github.com/stratexio/perms/whitespace"
)

//NamespaceSeperator seperates node namespaces
const NamespaceSeperator = "."

//WildcardSelector matches any namespace
const WildcardSelector = "*"

//NegateSignifier signifies a negation
const NegateSignifier = '-'

//common errors
var (
	ErrWhitespace  = errors.New("an illegal whitespace is present")
	ErrEmptyString = errors.New("node string is empty")
)

//Node contains a single permission node
type Node struct {
	Namespaces []string
	Negate     bool
}

//ParseNode parses a permission node
func ParseNode(raw string) (Node, error) {
	if whitespace.Contains(raw) {
		return Node{}, ErrWhitespace
	}
	tokens := strings.Split(raw, NamespaceSeperator)

	if len(tokens) == 0 {
		return Node{}, ErrEmptyString
	}

	if len(tokens[0]) == 0 {
		return Node{}, ErrEmptyString
	}

	var negate bool

	if tokens[0][0] == NegateSignifier {
		negate = true
		tokens[0] = tokens[0][1:]
	}

	return Node{
		Namespaces: tokens,
		Negate:     negate,
	}, nil
}

//Check checks if n permits check and considers negation
func (n Node) Check(check Node) bool {
	matched := n.Check(check)

	if matched {
		if check.Negate && n.Negate {
			return true
		}
		if check.Negate {
			return false
		}
		return matched
	}

	return false
}

//Match checks if a node matches another node.
//it is unaware of negation.
func (n Node) Match(check Node) bool {
	//If either node is empty the check fails
	if len(n.Namespaces) == 0 || len(check.Namespaces) == 0 {
		return false
	}
	var lastWildcard bool

	for i, namespace := range check.Namespaces {
		if len(n.Namespaces) == i {
			return lastWildcard
		}

		if n.Namespaces[i] == WildcardSelector {
			lastWildcard = true
			continue
		} else {
			lastWildcard = false
		}

		if namespace != n.Namespaces[i] {
			return false
		}
	}

	return !(len(check.Namespaces) < len(n.Namespaces))
}
