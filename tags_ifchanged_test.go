package pongo2_test

import (
	"testing"

	"github.com/flosch/pongo2/v6"
)

// TestIfchangedNoElseDoesNotPanic guards against a nil-pointer dereference
// when {% ifchanged %} has no {% else %} branch and the watched value does
// not change between iterations. The unchanged branch must produce no
// output instead of panicking on a nil elseWrapper.
func TestIfchangedNoElseDoesNotPanic(t *testing.T) {
	tpl, err := pongo2.FromString(`{% for v in items %}{% ifchanged v %}{{ v }}{% endifchanged %}{% endfor %}`)
	if err != nil {
		t.Fatal(err)
	}

	got, err := tpl.Execute(pongo2.Context{"items": []string{"a", "a", "b"}})
	if err != nil {
		t.Fatalf("execute: %v", err)
	}
	if got != "ab" {
		t.Fatalf("got %q want %q", got, "ab")
	}
}
