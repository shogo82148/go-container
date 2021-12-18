package optional

import "testing"

func TestOptional(t *testing.T) {
	valid := New(42)
	if !valid.Valid() {
		t.Error("want valid, but not")
	}

	invalid := None[int]()
	if invalid.Valid() {
		t.Error("want invalid, but valid")
	}
}
