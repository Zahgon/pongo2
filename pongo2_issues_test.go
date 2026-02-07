package pongo2_test

import (
	"sync"
	"testing"

	"github.com/flosch/pongo2/v6"
)

func TestIssue151(t *testing.T) {
	tpl, err := pongo2.FromString("{{ mydict.51232_3 }}{{ 12345_123}}{{ 995189baz }}")
	if err != nil {
		t.Fatal(err)
	}

	str, err := tpl.Execute(pongo2.Context{
		"mydict": map[string]string{
			"51232_3": "foo",
		},
		"12345_123": "bar",
		"995189baz": "baz",
	})
	if err != nil {
		t.Fatal(err)
	}

	if str != "foobarbaz" {
		t.Fatalf("Expected output 'foobarbaz', but got '%s'.", str)
	}
}

func TestBugCycleSharedState(t *testing.T) {
	// Bug: tagCycleNode.idx was stored on the AST node, which is shared
	// across all concurrent executions of the same parsed template.
	// This caused two problems:
	// 1. Data race: concurrent idx++ without synchronization
	// 2. Semantic bug: one execution's cycle position affected another,
	//    producing wrong output (e.g., "bcabca" instead of "abcabc")
	//
	// Each template execution must have independent cycle state.

	tpl, err := pongo2.FromString(`{% for i in items %}{% cycle "a" "b" "c" %}{% endfor %}`)
	if err != nil {
		t.Fatalf("failed to parse template: %v", err)
	}

	ctx := pongo2.Context{"items": []int{1, 2, 3, 4, 5, 6}}
	const expected = "abcabc"

	// First: verify sequential executions produce consistent results.
	// With shared state on the node, the second execution would start
	// where the first left off (idx=6), producing wrong output.
	for i := 0; i < 5; i++ {
		result, err := tpl.Execute(ctx)
		if err != nil {
			t.Fatalf("execution %d: unexpected error: %v", i, err)
		}
		if result != expected {
			t.Errorf("execution %d: got %q, want %q", i, result, expected)
		}
	}

	// Second: verify concurrent executions are also correct.
	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			result, err := tpl.Execute(ctx)
			if err != nil {
				t.Errorf("concurrent: unexpected error: %v", err)
				return
			}
			if result != expected {
				t.Errorf("concurrent: got %q, want %q", result, expected)
			}
		}()
	}
	wg.Wait()
}
