package perms

import "testing"

func TestParsePConf(t *testing.T) {
	byt := []byte(`
	{
    "groups": {
        "project_lead": {
            "nodes": [
                "analytics.*"
            ]
        },
        "manager": {
            "parents": [
                "project_lead"
            ],
            "nodes": [
                "projects.*"
            ]
        }
    },
    "users": {
        "ammar": {
            "groups": [
                "manager"
            ],
            "nodes": [
                "-projects.*.chat.moderate"
            ]
        }
    }
}
`)

	pconf, err := ParsePConf(byt)
	if err != nil {
		t.Fatalf("Failed to parse pconf: %v", err)
	}

	//fmt.Printf("%v\n", pconf)
	if pconf.Groups["manager"].Parents[0] != "project_lead" {
		t.Fatalf("Parse failed: pconf: %v", pconf)
	}
}
