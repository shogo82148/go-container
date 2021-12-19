package slices

import (
	"testing"

	"github.com/shogo82148/go-container/tuples"
)

func TestZip(t *testing.T) {
	in1 := []int{1, 2, 3, 4, 5}
	in2 := []string{"one", "two", "three", "four", "five"}
	got := Zip(in1, in2)
	want := []tuples.Tuple2[int, string]{
		{1, "one"},
		{2, "two"},
		{3, "three"},
		{4, "four"},
		{5, "five"},
	}

	if len(got) != len(want) {
		t.Fatalf("unexpected length: want %d, got %d", len(in1), len(got))
	}
	for i := range got {
		if got[i] != want[i] {
			t.Errorf("unexpected item: want got[%d] = %v, want[%d] = %v", i, got[i], i, want[i])
		}
	}
}

func TestZip_shortest(t *testing.T) {
	in1 := []int{1, 2, 3, 4, 5}
	in2 := []string{"one", "two", "three", "four"}
	got := Zip(in1, in2)
	want := []tuples.Tuple2[int, string]{
		{1, "one"},
		{2, "two"},
		{3, "three"},
		{4, "four"},
	}

	if len(got) != len(want) {
		t.Fatalf("unexpected length: want %d, got %d", len(in1), len(got))
	}
	for i := range got {
		if got[i] != want[i] {
			t.Errorf("unexpected item: want got[%d] = %v, want[%d] = %v", i, got[i], i, want[i])
		}
	}
}
