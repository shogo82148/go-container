package tuples

import "testing"

func TestTuple(t *testing.T) {
	pair := Tuple2[int, string]{42, "foobar"}
	if pair.V1 != 42 {
		t.Errorf("unexpected value: want %d, got %d", 42, pair.V1)
	}
	if pair.V2 != "foobar" {
		t.Errorf("unexpected value: want %q, got %q", "foobar", pair.V2)
	}
}
