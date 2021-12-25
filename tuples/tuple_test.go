package tuples

import "testing"

func TestTuple(t *testing.T) {
	pair := New2(42, "foobar")
	if pair.V1 != 42 {
		t.Errorf("unexpected value: want %d, got %d", 42, pair.V1)
	}
	if pair.V2 != "foobar" {
		t.Errorf("unexpected value: want %q, got %q", "foobar", pair.V2)
	}

	// compile test
	awesomeFunc(pair.Get())
}

func awesomeFunc(int, string) {}

func TestString(t *testing.T) {
	pair := New2(42, "foobar")
	got := pair.String()
	want := "(42, foobar)"
	if got != want {
		t.Errorf("want %q, got %q", want, got)
	}
}
