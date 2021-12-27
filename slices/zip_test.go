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
		{V1: 1, V2: "one"},
		{V1: 2, V2: "two"},
		{V1: 3, V2: "three"},
		{V1: 4, V2: "four"},
		{V1: 5, V2: "five"},
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

func TestZip2(t *testing.T) {
	in1 := []int{1, 2, 3, 4, 5}
	in2 := []string{"one", "two", "three", "four", "five"}
	got := Zip2(in1, in2)
	want := []tuples.Tuple2[int, string]{
		{V1: 1, V2: "one"},
		{V1: 2, V2: "two"},
		{V1: 3, V2: "three"},
		{V1: 4, V2: "four"},
		{V1: 5, V2: "five"},
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
		{V1: 1, V2: "one"},
		{V1: 2, V2: "two"},
		{V1: 3, V2: "three"},
		{V1: 4, V2: "four"},
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

func TestZip2_shortest(t *testing.T) {
	in1 := []int{1, 2, 3, 4, 5}
	in2 := []string{"one", "two", "three", "four"}
	got := Zip2(in1, in2)
	want := []tuples.Tuple2[int, string]{
		{V1: 1, V2: "one"},
		{V1: 2, V2: "two"},
		{V1: 3, V2: "three"},
		{V1: 4, V2: "four"},
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

func TestZip_underlying(t *testing.T) {
	type IntStrings []tuples.Tuple2[int, string]
	in1 := []int{1, 2, 3, 4, 5}
	in2 := []string{"one", "two", "three", "four", "five"}
	got := Zip[IntStrings](in1, in2)
	want := IntStrings{
		{V1: 1, V2: "one"},
		{V1: 2, V2: "two"},
		{V1: 3, V2: "three"},
		{V1: 4, V2: "four"},
		{V1: 5, V2: "five"},
	}

	if _, ok := (interface{})(got).(IntStrings); !ok {
		t.Errorf("want IntStrings type, got %T", got)
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

func TestZip2_underlying(t *testing.T) {
	type IntStrings []tuples.Tuple2[int, string]
	in1 := []int{1, 2, 3, 4, 5}
	in2 := []string{"one", "two", "three", "four", "five"}
	got := Zip2[IntStrings](in1, in2)
	want := IntStrings{
		{V1: 1, V2: "one"},
		{V1: 2, V2: "two"},
		{V1: 3, V2: "three"},
		{V1: 4, V2: "four"},
		{V1: 5, V2: "five"},
	}

	if _, ok := (interface{})(got).(IntStrings); !ok {
		t.Errorf("want IntStrings type, got %T", got)
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
