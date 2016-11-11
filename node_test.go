package permissions

import (
	"reflect"
	"testing"
)

func TestParseNode(t *testing.T) {
	type args struct {
		raw string
	}
	tests := []struct {
		name    string
		args    args
		want    Node
		wantErr bool
	}{
		{"simple", args{"projects.manage"}, Node{Namespaces: []string{"projects", "manage"}}, false},
		{"simple", args{"projects.manage.*"}, Node{Namespaces: []string{"projects", "manage", "*"}}, false},
		{"negate", args{"-projects.manage.*"}, Node{Namespaces: []string{"projects", "manage", "*"}, Negate: true}, false},
		{"whitespace", args{"- projects.manage.*"}, Node{}, true},
		{"empty", args{"..*"}, Node{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseNode(tt.args.raw)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseNode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNode_Check(t *testing.T) {
	type args struct {
		check Node
	}
	tests := []struct {
		name string
		n    Node
		args args
		want bool
	}{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.Check(tt.args.check); got != tt.want {
				t.Errorf("Node.Check() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNode_Match(t *testing.T) {
	type fields struct {
		Namespaces []string
		Negate     bool
	}
	type args struct {
		check Node
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := Node{
				Namespaces: tt.fields.Namespaces,
				Negate:     tt.fields.Negate,
			}
			if got := n.Match(tt.args.check); got != tt.want {
				t.Errorf("Node.Match() = %v, want %v", got, tt.want)
			}
		})
	}
}
