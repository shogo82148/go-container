package sets

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/shogo82148/go-container/tuples"
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
			a:    New("a"),
			b:    New[string](),
			want: false,
		},
		{
			a:    New[string](),
			b:    New("a"),
			want: false,
		},
		{
			a:    New("a"),
			b:    New("a"),
			want: true,
		},
		{
			a:    New("a"),
			b:    New("b"),
			want: false,
		},
		{
			a:    New("a", "b", "c"),
			b:    New("c", "b", "a"),
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

func TestClone(t *testing.T) {
	set := New(1, 2, 3)
	got := set.Clone()

	// update the original set.
	set.Add(4)

	// the update doesn't effect the new set.
	want := New(1, 2, 3)
	if !got.Equal(want) {
		t.Errorf("want %v, got %v", want, got)
	}
}

func BenchmarkClone(b *testing.B) {
	set := New[int]()
	for i := 0; i < 10000; i++ {
		set.Add(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		runtime.KeepAlive(set.Clone())
	}
}

func TestUnion(t *testing.T) {
	cases := []struct {
		a    Set[int]
		b    Set[int]
		want Set[int]
	}{
		{
			a:    New[int](),
			b:    New[int](),
			want: New[int](),
		},
		{
			a:    New(1, 2, 3),
			b:    New[int](),
			want: New(1, 2, 3),
		},
		{
			a:    New[int](),
			b:    New(1, 2, 3),
			want: New(1, 2, 3),
		},
		{
			a:    New(1, 2, 3),
			b:    New(1, 2, 3),
			want: New(1, 2, 3),
		},
		{
			a:    New(1, 2, 3),
			b:    New(2, 3, 4),
			want: New(1, 2, 3, 4),
		},
		{
			a:    New(1, 2, 3),
			b:    New(6, 7, 8),
			want: New(1, 2, 3, 6, 7, 8),
		},
	}

	for _, tt := range cases {
		got := tt.a.Union(tt.b)
		if !got.Equal(tt.want) {
			t.Errorf("want %v, got %v", tt.want, got)
		}
	}
}

func BenchmarkUnion(b *testing.B) {
	b.Run("max result", func(b *testing.B) {
		s := New[int]()
		t := New[int]()
		for i := 0; i < 10000; i++ {
			s.Add(i)
			t.Add(-i)
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			runtime.KeepAlive(s.Union(t))
		}
	})

	b.Run("min result", func(b *testing.B) {
		s := New[int]()
		t := New[int]()
		for i := 0; i < 10000; i++ {
			s.Add(i)
			t.Add(i)
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			runtime.KeepAlive(s.Union(t))
		}
	})
}

func TestIntersection(t *testing.T) {
	cases := []struct {
		a    Set[int]
		b    Set[int]
		want Set[int]
	}{
		{
			a:    New[int](),
			b:    New[int](),
			want: New[int](),
		},
		{
			a:    New(1, 2, 3),
			b:    New[int](),
			want: New[int](),
		},
		{
			a:    New[int](),
			b:    New(1, 2, 3),
			want: New[int](),
		},
		{
			a:    New(1, 2, 3),
			b:    New(1, 2, 3),
			want: New(1, 2, 3),
		},
		{
			a:    New(1, 2, 3),
			b:    New(2, 3, 4),
			want: New(2, 3),
		},
		{
			a:    New(1, 2, 3),
			b:    New(6, 7, 8),
			want: New[int](),
		},
	}

	for _, tt := range cases {
		got := tt.a.Intersection(tt.b)
		if !got.Equal(tt.want) {
			t.Errorf("want %v, got %v", tt.want, got)
		}
	}
}

func BenchmarkIntersection(b *testing.B) {
	b.Run("max result", func(b *testing.B) {
		s := New[int]()
		t := New[int]()
		for i := 0; i < 10000; i++ {
			s.Add(i)
			t.Add(i)
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			runtime.KeepAlive(s.Intersection(t))
		}
	})

	b.Run("min result", func(b *testing.B) {
		s := New[int]()
		t := New[int]()
		for i := 0; i < 10000; i++ {
			s.Add(i)
			t.Add(-i)
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			runtime.KeepAlive(s.Intersection(t))
		}
	})
}

func TestDifference(t *testing.T) {
	cases := []struct {
		a    Set[int]
		b    Set[int]
		want Set[int]
	}{
		{
			a:    New[int](),
			b:    New[int](),
			want: New[int](),
		},
		{
			a:    New(1, 2, 3),
			b:    New[int](),
			want: New(1, 2, 3),
		},
		{
			a:    New[int](),
			b:    New(1, 2, 3),
			want: New[int](),
		},
		{
			a:    New(1, 2, 3),
			b:    New(1, 2, 3),
			want: New[int](),
		},
		{
			a:    New(1, 2, 3),
			b:    New(2, 3, 4),
			want: New(1),
		},
		{
			a:    New(1, 2, 3),
			b:    New(6, 7, 8),
			want: New(1, 2, 3),
		},
	}

	for _, tt := range cases {
		got := tt.a.Difference(tt.b)
		if !got.Equal(tt.want) {
			t.Errorf("want %v, got %v", tt.want, got)
		}
	}
}

func BenchmarkDifference(b *testing.B) {
	b.Run("max result", func(b *testing.B) {
		s := New[int]()
		t := New[int]()
		for i := 0; i < 10000; i++ {
			s.Add(i)
			t.Add(-i)
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			runtime.KeepAlive(s.Difference(t))
		}
	})

	b.Run("min result", func(b *testing.B) {
		s := New[int]()
		t := New[int]()
		for i := 0; i < 10000; i++ {
			s.Add(i)
			t.Add(i)
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			runtime.KeepAlive(s.Difference(t))
		}
	})
}

func TestSymmetricDifference(t *testing.T) {
	cases := []struct {
		a    Set[int]
		b    Set[int]
		want Set[int]
	}{
		{
			a:    New[int](),
			b:    New[int](),
			want: New[int](),
		},
		{
			a:    New(1, 2, 3),
			b:    New[int](),
			want: New(1, 2, 3),
		},
		{
			a:    New[int](),
			b:    New(1, 2, 3),
			want: New(1, 2, 3),
		},
		{
			a:    New(1, 2, 3),
			b:    New(1, 2, 3),
			want: New[int](),
		},
		{
			a:    New(1, 2, 3),
			b:    New(2, 3, 4),
			want: New(1, 4),
		},
		{
			a:    New(1, 2, 3),
			b:    New(6, 7, 8),
			want: New(1, 2, 3, 6, 7, 8),
		},
	}

	for _, tt := range cases {
		got := tt.a.SymmetricDifference(tt.b)
		if !got.Equal(tt.want) {
			t.Errorf("want %v, got %v", tt.want, got)
		}
	}
}

func BenchmarkSymmetricDifference(b *testing.B) {
	b.Run("max result", func(b *testing.B) {
		s := New[int]()
		t := New[int]()
		for i := 0; i < 10000; i++ {
			s.Add(i)
			t.Add(-i)
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			runtime.KeepAlive(s.SymmetricDifference(t))
		}
	})

	b.Run("min result", func(b *testing.B) {
		s := New[int]()
		t := New[int]()
		for i := 0; i < 10000; i++ {
			s.Add(i)
			t.Add(i)
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			runtime.KeepAlive(s.SymmetricDifference(t))
		}
	})
}

func TestIsSubset(t *testing.T) {
	cases := []struct {
		a                Set[int]
		b                Set[int]
		isSubset         bool
		isProperSubset   bool
		isSuperset       bool
		isProperSuperset bool
	}{
		{
			a:                New[int](),
			b:                New[int](),
			isSubset:         true,
			isProperSubset:   false,
			isSuperset:       true,
			isProperSuperset: false,
		},
		{
			a:                New[int](),
			b:                New(1, 2, 3),
			isSubset:         true,
			isProperSubset:   true,
			isSuperset:       false,
			isProperSuperset: false,
		},
		{
			a:                New(1),
			b:                New(1, 2, 3),
			isSubset:         true,
			isProperSubset:   true,
			isSuperset:       false,
			isProperSuperset: false,
		},
		{
			a:                New(1, 2, 3),
			b:                New(1, 2, 3),
			isSubset:         true,
			isProperSubset:   false,
			isSuperset:       true,
			isProperSuperset: false,
		},
		{
			a:                New(1, 2, 3, 4),
			b:                New(1, 2, 3),
			isSubset:         false,
			isProperSubset:   false,
			isSuperset:       true,
			isProperSuperset: true,
		},
		{
			a:                New(2, 3, 4),
			b:                New(1, 2, 3),
			isSubset:         false,
			isProperSubset:   false,
			isSuperset:       false,
			isProperSuperset: false,
		},
	}

	for _, tt := range cases {
		isSubset := tt.a.IsSubset(tt.b)
		if isSubset != tt.isSubset {
			t.Errorf("%v.IsSubset(%v): want %t, got %t", tt.a, tt.b, tt.isSubset, isSubset)
		}
		isProperSubset := tt.a.IsProperSubset(tt.b)
		if isProperSubset != tt.isProperSubset {
			t.Errorf("%v.IsProperSubset(%v): want %t, got %t", tt.a, tt.b, tt.isProperSubset, isProperSubset)
		}
		isSuperset := tt.a.IsSuperset(tt.b)
		if isSuperset != tt.isSuperset {
			t.Errorf("%v.IsSuperset(%v): want %t, got %t", tt.a, tt.b, tt.isSuperset, isSuperset)
		}
		isProperSuperset := tt.a.IsProperSuperset(tt.b)
		if isProperSuperset != tt.isProperSuperset {
			t.Errorf("%v.IsProperSuperset(%v): want %t, got %t", tt.a, tt.b, tt.isProperSuperset, isProperSuperset)
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

func TestProduct(t *testing.T) {
	a := New(1, 2, 3)
	b := New("a", "b", "c")
	got := Product(a, b)
	want := New(
		tuples.New2(1, "a"), tuples.New2(1, "b"), tuples.New2(1, "c"),
		tuples.New2(2, "a"), tuples.New2(2, "b"), tuples.New2(2, "c"),
		tuples.New2(3, "a"), tuples.New2(3, "b"), tuples.New2(3, "c"),
	)
	if !got.Equal(want) {
		t.Errorf("want %v, got %v", want, got)
	}
}
