package list

import (
	"testing"

	"github.com/shogo82148/go-container/tuples"
)

func TestZip(t *testing.T) {
	var in1 List[int]
	for _, v := range []int{1, 2, 3, 4, 5} {
		in1.PushBack(v)
	}

	var in2 List[string]
	for _, v := range []string{"one", "two", "three", "four", "five"} {
		in2.PushBack(v)
	}
	got := Zip(in1, in2)
	want := []tuples.Tuple2[int, string]{
		{V1: 1, V2: "one"},
		{V1: 2, V2: "two"},
		{V1: 3, V2: "three"},
		{V1: 4, V2: "four"},
		{V1: 5, V2: "five"},
	}

	if got.Len() != len(want) {
		t.Fatalf("unexpected length: want %d, got %d", len(want), got.Len())
	}
	for i, e := 0, got.Front(); i < len(want) && e != nil; i, e = i+1, e.Next() {
		if e.Value != want[i] {
			t.Errorf("unexpected item: want got[%d] = %v, want[%d] = %v", i, e.Value, i, want[i])
		}
	}
}
