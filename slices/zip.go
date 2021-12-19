package slices

import (
	"fmt"

	"github.com/shogo82148/go-container/tuples"
)

// Zip returns a slice of 2-tuples.
// All slices must have same length. the lengths are different, Zip panics.
func Zip[T1, T2 any](s1 []T1, s2 []T2) []tuples.Tuple2[T1, T2] {
	l := len(s1)
	if len(s2) != l {
		panic(fmt.Errorf("s2 have different length from s1. len(s2) = %d, len(s1) = %d", len(s2), l))
	}
	ret := make([]tuples.Tuple2[T1, T2], l)
	for i := 0; i < l; i++ {
		ret[i] = tuples.Tuple2[T1, T2]{s1[i], s2[i]}
	}
	return ret
}
