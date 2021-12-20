package optional

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	called := false
	f := func(v int) string {
		called = true
		return fmt.Sprintf("<%d>", v)
	}

	got := Map(New(42), f)
	if !called {
		t.Error("the function should be called, but not")
	}
	if !got.Valid() {
		t.Error("want valid, but got invalid")
	}
	if got.Get() != "<42>" {
		t.Errorf("want %q, got %q", "<42>", got.Get())
	}

	called = false
	got = Map(None[int](), f)
	if called {
		t.Error("the function should not be called, but is called")
	}
	if got.Valid() {
		t.Error("want invalid, but got valid")
	}
}
