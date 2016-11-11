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
		{"empty", args{""}, Node{}, true},
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

func TestNode_Match(t *testing.T) {
	type args struct {
		check Node
	}
	tests := []struct {
		name string
		n    Node
		args args
		want bool
	}{
		{"simple", Node{Namespaces: []string{"projects", "webserver"}}, args{Node{Namespaces: []string{"projects", "webserver"}}}, true},
		{"simple", Node{Namespaces: []string{"projects", "webserver"}}, args{Node{Namespaces: []string{"projects", "frontend"}}}, false},

		{"wildcard", Node{Namespaces: []string{"projects", "*"}}, args{Node{Namespaces: []string{"projects", "frontend"}}}, true},
		{"wildcard", Node{Namespaces: []string{"projects", "*"}}, args{Node{Namespaces: []string{"billing", "frontend"}}}, false},

		{"middle_wildcard", Node{Namespaces: []string{"projects", "*", "chat"}}, args{Node{Namespaces: []string{"projects", "test"}}}, false},
		{"middle_wildcard", Node{Namespaces: []string{"projects", "*", "chat"}}, args{Node{Namespaces: []string{"projects", "test", "test"}}}, false},
		{"middle_wildcard", Node{Namespaces: []string{"projects", "*", "chat"}}, args{Node{Namespaces: []string{"projects", "test", "chat"}}}, true},

		{"supernode", Node{Namespaces: []string{"*"}}, args{Node{Namespaces: []string{"projects", "test", "chat"}}}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.Match(tt.args.check); got != tt.want {
				t.Errorf("Node.Check() = %v, want %v", got, tt.want)
			}
		})
	}
}
