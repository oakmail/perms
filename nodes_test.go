package perms

import (
	"reflect"
	"testing"
)

func TestParseNodes(t *testing.T) {
	type args struct {
		raw []byte
	}
	tests := []struct {
		name    string
		args    args
		want    Nodes
		wantErr bool
	}{
		{
			"simple",
			args{raw: []byte("project.test project.build")},
			Nodes{
				MustParseNode("project.test"),
				MustParseNode("project.build"),
			},
			false,
		},
		{
			"simple",
			args{raw: []byte("project.test\nproject.build")},
			Nodes{
				MustParseNode("project.test"),
				MustParseNode("project.build"),
			},
			false,
		},
		{
			"whitepower",
			args{raw: []byte("   project.test  \nproject.build ")},
			Nodes{
				MustParseNode("project.test"),
				MustParseNode("project.build"),
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseNodes(tt.args.raw)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseNodes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNodes_String(t *testing.T) {
	tests := []struct {
		name string
		ns   Nodes
		want string
	}{
		{
			"simple",
			Nodes{
				MustParseNode("project.build"),
				MustParseNode("project.test"),
			},
			"project.build\nproject.test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ns.String(); got != tt.want {
				t.Errorf("Nodes.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
