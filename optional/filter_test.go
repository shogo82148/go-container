package optional

import (
	"testing"
)

func TestFilter(t *testing.T) {
	cases := []struct {
		in    Optional[int]
		valid bool
	}{
		{
			in:    New(3),
			valid: false,
		},
		{
			in:    New(4),
			valid: true,
		},
		{
			in:    None[int](),
			valid: false,
		},
	}

	f := func(v int) bool {
		return v > 3
	}
	for _, tt := range cases {
		got := Filter(tt.in, f)
		if got.Valid() != tt.valid {
			t.Errorf("want %t, got %t", tt.valid, got.Valid())
		}
	}
}
