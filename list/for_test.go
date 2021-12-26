package list

import (
	"testing"
)

func TestFor(t *testing.T) {
	l := New[int]()
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	l.PushBack(4)

	got := []int{}
	f := func(e *Element[int]) error {
		got = append(got, e.Value)
		return nil
	}
	if err := l.For(f); err != nil {
		t.Fatal(err)
	}

	want := []int{1, 2, 3, 4}

	if len(got) != len(want) {
		t.Errorf("unexpected length: want %d, got %d", len(want), len(got))
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("unexpected value at %d: want %d, got %d", i, want[i], got[i])
		}
	}
}
