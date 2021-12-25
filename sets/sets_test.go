package sets

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	t.Run("empty set", func(t *testing.T) {
		set := New[string]()
		if set == nil {
			t.Error("want not nil, got nil")
		}
		if len(set) != 0 {
			t.Errorf("want %d, got %d", 0, len(set))
		}
	})

	t.Run("strings", func(t *testing.T) {
		set := New("a", "b")
		if set == nil {
			t.Error("want not nil, got nil")
		}
		if len(set) != 2 {
			t.Errorf("want %d, got %d", 2, len(set))
		}
		if !set.Contains("a") {
			t.Error("want contains \"a\", but not")
		}
		if !set.Contains("b") {
			t.Error("want contains \"b\", but not")
		}
		if set.Contains("c") {
			t.Error("want not contains \"c\", but it does")
		}
	})

	t.Run("ints", func(t *testing.T) {
		set := New(1, 2)
		if set == nil {
			t.Error("want not nil, got nil")
		}
		if len(set) != 2 {
			t.Errorf("want %d, got %d", 2, len(set))
		}
		if !set.Contains(1) {
			t.Error("want contains 1, but not")
		}
		if !set.Contains(2) {
			t.Error("want contains 2, but not")
		}
		if set.Contains(3) {
			t.Error("want not contains 3, but it does")
		}
	})
}

func TestAdd(t *testing.T) {
	set := New("a", "b")
	if len(set) != 2 {
		t.Errorf("want %d, got %d", 2, len(set))
	}
	if set.Contains("c") {
		t.Error("want not contains \"c\", but it does")
	}

	// add "c"
	set.Add("c")
	if len(set) != 3 {
		t.Errorf("want %d, got %d", 2, len(set))
	}
	if !set.Contains("c") {
		t.Error("want contains \"c\", but not")
	}

	// add "c" again
	set.Add("c")
	if len(set) != 3 {
		t.Errorf("want %d, got %d", 2, len(set))
	}
}

func TestRemove(t *testing.T) {
	set := New("a", "b")
	if len(set) != 2 {
		t.Errorf("want %d, got %d", 2, len(set))
	}
	if !set.Contains("a") {
		t.Error("want contains \"a\", but not")
	}

	// remove "a"
	set.Remove("a")
	if len(set) != 1 {
		t.Errorf("want %d, got %d", 1, len(set))
	}
	if set.Contains("a") {
		t.Error("want not contains \"a\", but it does")
	}

	// remove "a" again
	set.Remove("a")
	if len(set) != 1 {
		t.Errorf("want %d, got %d", 1, len(set))
	}
	if set.Contains("a") {
		t.Error("want not contains \"a\", but it does")
	}
}

func TestLen(t *testing.T) {
	set := New("a", "b")
	if set.Len() != 2 {
		t.Errorf("want %d, got %d", 2, set.Len())
	}
}

func TestEqual(t *testing.T) {
	cases := []struct {
		a    Set[string]
		b    Set[string]
		want bool
	}{
		{
			a:    New[string](),
			b:    New[string](),
			want: true,
		},
		{
			a:    New[string]("a"),
			b:    New[string](),
			want: false,
		},
		{
			a:    New[string](),
			b:    New[string]("a"),
			want: false,
		},
		{
			a:    New[string]("a"),
			b:    New[string]("a"),
			want: true,
		},
		{
			a:    New[string]("a"),
			b:    New[string]("b"),
			want: false,
		},
		{
			a:    New[string]("a", "b", "c"),
			b:    New[string]("c", "b", "a"),
			want: true,
		},
	}

	for _, tt := range cases {
		got := tt.a.Equal(tt.b)
		if got != tt.want {
			t.Errorf("want %t, got %t", tt.want, got)
		}
	}
}

func TestFor(t *testing.T) {
	set := New(1, 2, 3)
	got := New[string]()
	set.For(func(v int) error {
		got.Add(fmt.Sprintf("<%d>", v))
		return nil
	})
	want := New("<1>", "<2>", "<3>")
	if !got.Equal(want) {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestFilter(t *testing.T) {
	set := New(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	got := set.Filter(func(v int) bool {
		return v > 5
	})
	want := New(6, 7, 8, 9, 10)
	if !got.Equal(want) {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestMap(t *testing.T) {
	in := New(1, 2, 3)
	got := Map(in, func(v int) string {
		return fmt.Sprintf("<%d>", v)
	})
	want := New("<1>", "<2>", "<3>")
	if !got.Equal(want) {
		t.Errorf("want %v, got %v", want, got)
	}
}
