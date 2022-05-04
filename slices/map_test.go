package slices

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	in := []int{1, 2, 3, 4, 5, 6}
	f := func(a int) string {
		return fmt.Sprintf("<%d>", a)
	}
	got := Map[[]string](in, f)

	want := []string{"<1>", "<2>", "<3>", "<4>", "<5>", "<6>"}

	if len(got) != len(want) {
		t.Errorf("unexpected length: want %d, got %d", len(want), len(got))
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("unexpected value at %d: want %q, got %q", i, want[i], got[i])
		}
	}
}
