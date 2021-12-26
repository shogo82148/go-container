package list

import (
	"testing"
)

func TestFilter(t *testing.T) {
	l := New[int]()
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	l.PushBack(4)
	l.PushBack(5)
	l.PushBack(6)

	got := l.Filter(func(e *Element[int]) bool {
		return e.Value > 3
	})

	want := []int{4, 5, 6}
	if got.Len() != len(want) {
		t.Errorf("unexpected length: want %d, got %d", len(want), got.Len())
	}
	for i, e := 0, got.Front(); e != nil; i, e = i+1, e.Next() {
		if e.Value != want[i] {
			t.Errorf("unexpected value, want got[%d] = %d, got %d", i, want[i], e.Value)
		}
	}
}
