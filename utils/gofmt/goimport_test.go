package gofmt

import "testing"

func TestGofmtMain(t *testing.T) {
	path := "/Users/miraculous/application_library/go_path/src/code.byted.org/comments/comments_build_tools/gofmt/goimport.go"
	t.Log(GofmtMain(path))
}
