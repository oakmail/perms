package perms

import (
	"bytes"
	"strings"

	"github.com/oakmail/perms/whitespace"
	"github.com/pkg/errors"
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

//MustParseNode panics if an error occurs while parsing the node
func MustParseNode(raw string) Node {
	node, err := ParseNode(raw)
	if err != nil {
		panic(err)
	}
	return node
}

//Match checks if a node matches another node.
//it is unaware of negation.
func (n Node) Match(check Node) bool {
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

//String returns the string representation of the node
func (n Node) String() string {
	buf := new(bytes.Buffer)
	if n.Negate {
		buf.WriteByte('-')
	}
	for i, namespace := range n.Namespaces {
		buf.WriteString(namespace)
		if i != (len(n.Namespaces) - 1) {
			buf.WriteByte('.')
		}
	}
	return buf.String()
}
