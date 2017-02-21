package perms

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"

	"github.com/trtlio/sqltypes"

	"github.com/trtlio/perms/whitespace"
)

//Nodes is a list of nodes
type Nodes []Node

//ParseNodes parses a whitespace delimited list of nodes
func ParseNodes(raw []byte) (Nodes, error) {
	bufReader := bytes.NewReader(raw)
	nodes := make(Nodes, 0, 10)
	lastNodeText := new(bytes.Buffer)

	flush := func() error {
		if lastNodeText.Len() > 0 {
			node, err := ParseNode(lastNodeText.String())
			if err != nil {
				return err
			}
			nodes = append(nodes, node)
		}
		lastNodeText.Reset()
		return nil
	}

	for {
		r, _, err := bufReader.ReadRune()
		if err != nil {
			return nodes, flush()
		}
		if whitespace.Is(r) {
			if err := flush(); err != nil {
				return nil, err
			}
			continue
		}
		lastNodeText.WriteRune(r)
	}
}

//MustParseNodes parses raw or panics
func MustParseNodes(raw []byte) Nodes {
	n, err := ParseNodes(raw)
	if err != nil {
		panic(err)
	}
	return n
}

//Check checks for a permission with ns
func (ns Nodes) Check(check Node) (matched bool, negated bool) {
	for _, node := range ns {
		if node.Match(check) {
			matched = true
			if node.Negate {
				return true, true
			}
		}
	}
	return matched, false
}

//String returns a string representation of n
func (ns Nodes) String() string {
	if ns == nil {
		return ""
	}
	buf := new(bytes.Buffer)
	for k, n := range ns {
		buf.WriteString(n.String())
		if k != (len(ns) - 1) {
			buf.WriteByte(' ')
		}
	}
	return buf.String()
}

//Strings returns ns as a slice of it's individual strings
func (ns Nodes) Strings() []string {
	if ns == nil {
		return []string{}
	}
	strs := make([]string, len(ns))
	for i, node := range ns {
		strs[i] = node.String()
	}
	return strs
}

// Scan implements the SQL Scanner interface
func (ns *Nodes) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	var input []byte
	if err := sqltypes.ConvertAssign(&input, value); err != nil {
		return err
	}

	input = bytes.Replace(input, []byte("\\n"), []byte("\n"), -1)

	nn, err := ParseNodes(input)
	if err != nil {
		return err
	}

	*ns = nn
	return nil
}

// Value implements the SQL driver Valuer interface
func (ns Nodes) Value() (driver.Value, error) {
	return ns.String(), nil
}

// MarshalJSON implements the JSON marshaller interface
func (ns Nodes) MarshalJSON() ([]byte, error) {
	return json.Marshal(ns.String())
}

// UnmarshalJSON implements the JSON unmarshaller interface
func (ns Nodes) UnmarshalJSON(b []byte) error {
	var data string
	if err := json.Unmarshal(b, &data); err != nil {
		return err
	}

	nodes, err := ParseNodes([]byte(data))
	if err != nil {
		return err
	}

	ns = nodes

	return nil
}
