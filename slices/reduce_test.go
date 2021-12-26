package slices

import "testing"

func TestReduce(t *testing.T) {
	t.Run("sum", func(t *testing.T) {
		in := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		got := Reduce(in, 0, func(lhs, rhs int) int {
			return lhs + rhs
		})
		if got != 55 {
			t.Errorf("unexpected result: want 55, got %d", got)
		}
	})

	t.Run("map", func(t *testing.T) {
		in := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		got := Reduce(in, make([]int, 0, len(in)), func(lhs []int, rhs int) []int {
			return append(lhs, rhs*2)
		})
		for i := range in {
			if got[i] != in[i]*2 {
				t.Errorf("unexpected value: want %d, got %d", in[i]*2, got[i])
			}
		}
	})
}
