package sets

import "testing"

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
