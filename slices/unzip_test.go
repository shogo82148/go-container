package slices

import (
	"testing"

	"github.com/shogo82148/go-container/tuples"
)

func TestUnzip(t *testing.T) {
	in := []tuples.Tuple2[int, string]{
		{V1: 1, V2: "one"},
		{V1: 2, V2: "two"},
		{V1: 3, V2: "three"},
		{V1: 4, V2: "four"},
		{V1: 5, V2: "five"},
	}
	want1 := []int{1, 2, 3, 4, 5}
	want2 := []string{"one", "two", "three", "four", "five"}
	got1, got2 := Unzip[[]int, []string](in)

	if len(got1) != len(want1) {
		t.Fatalf("unexpected length: want %d, got %d", len(want1), len(got1))
	}
	if len(got2) != len(want2) {
		t.Fatalf("unexpected length: want %d, got %d", len(want2), len(got2))
	}
	for i := range got1 {
		if got1[i] != want1[i] {
			t.Errorf("unexpected item: want got1[%d] = %v, want1[%d] = %v", i, got1[i], i, want1[i])
		}
	}
	for i := range got2 {
		if got2[i] != want2[i] {
			t.Errorf("unexpected item: want got2[%d] = %v, want2[%d] = %v", i, got2[i], i, want2[i])
		}
	}
}

func TestUnzip2(t *testing.T) {
	in := []tuples.Tuple2[int, string]{
		{V1: 1, V2: "one"},
		{V1: 2, V2: "two"},
		{V1: 3, V2: "three"},
		{V1: 4, V2: "four"},
		{V1: 5, V2: "five"},
	}
	want1 := []int{1, 2, 3, 4, 5}
	want2 := []string{"one", "two", "three", "four", "five"}
	got1, got2 := Unzip2(in)

	if len(got1) != len(want1) {
		t.Fatalf("unexpected length: want %d, got %d", len(want1), len(got1))
	}
	if len(got2) != len(want2) {
		t.Fatalf("unexpected length: want %d, got %d", len(want2), len(got2))
	}
	for i := range got1 {
		if got1[i] != want1[i] {
			t.Errorf("unexpected item: want got1[%d] = %v, want1[%d] = %v", i, got1[i], i, want1[i])
		}
	}
	for i := range got2 {
		if got2[i] != want2[i] {
			t.Errorf("unexpected item: want got2[%d] = %v, want2[%d] = %v", i, got2[i], i, want2[i])
		}
	}
}
