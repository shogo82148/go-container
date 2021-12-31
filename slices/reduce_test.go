package slices

import "testing"

func TestReduce(t *testing.T) {
	t.Run("sum", func(t *testing.T) {
		in := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		got := Reduce(in, func(lhs, rhs int) int {
			return lhs + rhs
		})
		if got != 55 {
			t.Errorf("unexpected result: want 55, got %d", got)
		}
	})
}
