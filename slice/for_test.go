package slice

import (
	"fmt"
	"testing"
)

func TestFor(t *testing.T) {
	in := []int{1, 2, 3, 4, 5, 6}
	got := []string{}
	f := func(i, v int) error {
		got = append(got, fmt.Sprintf("%d: <%d>", i, v))
		return nil
	}
	if err := For(in, f); err != nil {
		t.Fatal(err)
	}

	want := []string{"0: <1>", "1: <2>", "2: <3>", "3: <4>", "4: <5>", "5: <6>"}

	if len(got) != len(want) {
		t.Errorf("unexpected length: want %d, got %d", len(want), len(got))
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("unexpected value at %d: want %q, got %q", i, want[i], got[i])
		}
	}
}
