package pongo2_test

import (
	"sync"
	"testing"

	"github.com/flosch/pongo2/v6"
)

// TestBanFilterRaceWithFromString exercises the documented invariant
// that BanFilter / BanTag must be called before any template is added
// to the set. Even when application code follows that rule, concurrent
// startup goroutines (configuration vs. template warmup) can race on
// the shared firstTemplateCreated flag and bannedFilters map. Under
// -race this test detects the race; without -race the same data race
// can produce a torn write or a missed ban.
func TestBanFilterRaceWithFromString(t *testing.T) {
	set := pongo2.NewSet("ban-race", pongo2.MustNewLocalFileSystemLoader(""))

	const goroutines = 8

	var wg sync.WaitGroup
	wg.Add(goroutines * 2)

	for i := 0; i < goroutines; i++ {
		go func() {
			defer wg.Done()
			// Try to ban a real filter; ignore the "already banned" /
			// "after first template" errors — we only care about race-free
			// access.
			_ = set.BanFilter("safe")
		}()
		go func() {
			defer wg.Done()
			_, _ = set.FromString("{{ x }}")
		}()
	}

	wg.Wait()
}
