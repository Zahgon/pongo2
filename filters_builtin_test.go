package pongo2_test

import (
	"testing"

	"github.com/flosch/pongo2/v6"
)

// TestRemovetagsRegexMetacharSafe ensures removetags does not panic when
// the tag-name parameter contains regex metacharacters. Such input must
// be treated as a literal tag name, not an arbitrary regular expression.
func TestRemovetagsRegexMetacharSafe(t *testing.T) {
	tpl, err := pongo2.FromString(`{{ "<a>x</a>"|removetags:bad|safe }}`)
	if err != nil {
		t.Fatal(err)
	}

	got, err := tpl.Execute(pongo2.Context{"bad": "["})
	if err != nil {
		t.Fatalf("execute: %v", err)
	}
	if got != "<a>x</a>" {
		t.Fatalf("got %q want %q", got, "<a>x</a>")
	}
}
