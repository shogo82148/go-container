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

func TestGetOr(t *testing.T) {
	got := New(42).GetOr(46)
	if got != 42 {
		t.Errorf("want 42, got %d", got)
	}

	got = None[int]().GetOr(46)
	if got != 46 {
		t.Errorf("want 46, got %d", got)
	}
}

func TestGetOrFunc(t *testing.T) {
	called := false
	f := func() int {
		called = true
		return 46
	}

	got := New(42).GetOrFunc(f)
	if got != 42 {
		t.Errorf("want 42, got %d", got)
	}
	if called {
		t.Error("the function should be not called, but is called")
	}

	called = false
	got = None[int]().GetOrFunc(f)
	if got != 46 {
		t.Errorf("want 46, got %d", got)
	}
	if !called {
		t.Error("the function should be called, but is not called")
	}
}
