package optional

import "testing"

var _ Optional[int] = Some[int]{}
var _ Optional[int] = None[int]{}

func TestNewSome(t *testing.T) {
	var v Optional[int] = NewSome[int](42)

	if !v.Valid() {
		t.Error("want valid, got invalid")
	}

	if got := v.GetOr(46); got != 42 {
		t.Errorf("want %d, got %d", 42, got)
	}

	if got := v.GetOrFunc(func() int { return 46 }); got != 42 {
		t.Errorf("want %d, got %d", 42, got)
	}

	switch v := v.(type) {
	case Some[int]:
		if v.Get() != 42 {
			t.Errorf("want %d, got %d", 42, v.Get())
		}
	case None[int]:
		t.Fatal("never reach")
	}
}

func TestNewNone(t *testing.T) {
	var v Optional[int] = NewNone[int]()

	if v.Valid() {
		t.Error("want invalid, got valid")
	}

	if got := v.GetOr(46); got != 46 {
		t.Errorf("want %d, got %d", 46, got)
	}

	if got := v.GetOrFunc(func() int { return 46 }); got != 46 {
		t.Errorf("want %d, got %d", 46, got)
	}

	switch v.(type) {
	case Some[int]:
		t.Fatal("never reach")
	case None[int]:
	}
}
