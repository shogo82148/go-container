package optional

import "testing"

func TestOptional(t *testing.T) {
	valid := New(42)
	if !valid.Valid() {
		t.Error("want valid, but not")
	}
	if valid.Get() != 42 {
		t.Errorf("want %d, got %d", 42, valid.Get())
	}

	invalid := None[int]()
	if invalid.Valid() {
		t.Error("want invalid, but valid")
	}

	paniced := false
	func() {
		defer func() {
			v := recover()
			paniced = v != nil
		}()
		invalid.Get()
	}()
	if !paniced {
		t.Error("should panic, but not")
	}
}
