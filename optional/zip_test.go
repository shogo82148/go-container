package optional

import (
	"testing"

	"github.com/shogo82148/go-container/tuples"
)

func TestZip(t *testing.T) {
	cases := []struct {
		in1  Optional[int]
		in2  Optional[string]
		want Optional[tuples.Tuple2[int, string]]
	}{
		{
			in1:  New(1),
			in2:  New("one"),
			want: New(tuples.New2(1, "one")),
		},
		{
			in1:  New(1),
			in2:  None[string](),
			want: None[tuples.Tuple2[int, string]](),
		},
		{
			in1:  None[int](),
			in2:  New("one"),
			want: None[tuples.Tuple2[int, string]](),
		},
		{
			in1:  None[int](),
			in2:  None[string](),
			want: None[tuples.Tuple2[int, string]](),
		},
	}

	for _, tt := range cases {
		got := Zip(tt.in1, tt.in2)
		if got != tt.want {
			t.Errorf("want %v, got %v", tt.want, got)
		}
	}
}

func TestZip2(t *testing.T) {
	cases := []struct {
		in1  Optional[int]
		in2  Optional[string]
		want Optional[tuples.Tuple2[int, string]]
	}{
		{
			in1:  New(1),
			in2:  New("one"),
			want: New(tuples.New2(1, "one")),
		},
		{
			in1:  New(1),
			in2:  None[string](),
			want: None[tuples.Tuple2[int, string]](),
		},
		{
			in1:  None[int](),
			in2:  New("one"),
			want: None[tuples.Tuple2[int, string]](),
		},
		{
			in1:  None[int](),
			in2:  None[string](),
			want: None[tuples.Tuple2[int, string]](),
		},
	}

	for _, tt := range cases {
		got := Zip2(tt.in1, tt.in2)
		if got != tt.want {
			t.Errorf("want %v, got %v", tt.want, got)
		}
	}
}
