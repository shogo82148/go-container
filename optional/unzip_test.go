package optional

import (
	"testing"

	"github.com/shogo82148/go-container/tuples"
)

func TestUnzip2(t *testing.T) {
	cases := []struct {
		in    Optional[tuples.Tuple2[int, string]]
		want1 Optional[int]
		want2 Optional[string]
	}{
		{
			in:    New(tuples.Tuple2[int, string]{1, "one"}),
			want1: New(1),
			want2: New("one"),
		},
		{
			in:    None[tuples.Tuple2[int, string]](),
			want1: None[int](),
			want2: None[string](),
		},
	}

	for _, tt := range cases {
		got1, got2 := Unzip2(tt.in)
		if got1 != tt.want1 {
			t.Errorf("want %v, got %v", tt.want1, got1)
		}
		if got2 != tt.want2 {
			t.Errorf("want %v, got %v", tt.want2, got2)
		}
	}
}
