package pongo2_test

import (
	"testing"

	"github.com/flosch/pongo2/v6"
)

// TestWidthratioMaxZero ensures widthratio renders an empty string when
// the max argument is zero, instead of producing garbage output from
// int(math.Ceil(NaN/Inf + 0.5)). This matches Django's behavior of
// returning an empty string on division-by-zero.
func TestWidthratioMaxZero(t *testing.T) {
	tpl, err := pongo2.FromString(`x{% widthratio cur max 100 %}y`)
	if err != nil {
		t.Fatal(err)
	}

	got, err := tpl.Execute(pongo2.Context{"cur": 5, "max": 0})
	if err != nil {
		t.Fatalf("execute: %v", err)
	}
	if got != "xy" {
		t.Fatalf("widthratio with max=0 produced %q, want %q", got, "xy")
	}
}
