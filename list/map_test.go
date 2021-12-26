package list

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	l := New[int]()
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)

	got := Map(l, func(v int) string {
		return fmt.Sprintf("<%d>", v)
	})

	want := []string{"<1>", "<2>", "<3>"}
	if got.Len() != len(want) {
		t.Errorf("unexpected length: want %d, got %d", len(want), got.Len())
	}
	for i, e := 0, got.Front(); e != nil; i, e = i+1, e.Next() {
		if e.Value != want[i] {
			t.Errorf("unexpected value, want got[%d] = %q, got %q", i, want[i], e.Value)
		}
	}
}
