package perms

import (
	"encoding/json"

	"github.com/pkg/errors"
)

//PConf contains a permissions config
type PConf struct {
	Groups map[string]struct {
		Parents []string `json:"parents"`
		Nodes   []string `json:"nodes"`
	} `json:"groups"`
	Users map[string]struct {
		Groups []string `json:"groups"`
		Nodes  []string `json:"nodes"`
	} `json:"users"`
}

//ParsePConf parses a pconf
func ParsePConf(byt []byte) (*PConf, error) {
	var pconf PConf
	if err := json.Unmarshal(byt, &pconf); err != nil {
		return nil, errors.Wrap(err, "failed to decode")
	}
	return &pconf, nil
}

//MustParsePConf parses the conf or panics trying
func MustParsePConf(byt []byte) *PConf {
	pcon, err := ParsePConf(byt)
	if err != nil {
		panic(err)
	}
	return pcon
}
