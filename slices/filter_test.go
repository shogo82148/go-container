package slices

import (
	"testing"
)

func TestFilter(t *testing.T) {
	in := []int{1, 2, 3, 4, 5, 6}
	f := func(v int) bool {
		return v > 3
	}
	got := Filter(in, f)

	want := []int{4, 5, 6}

	if len(got) != len(want) {
		t.Errorf("unexpected length: want %d, got %d", len(want), len(got))
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("unexpected value at %d: want %q, got %q", i, want[i], got[i])
		}
	}
}
