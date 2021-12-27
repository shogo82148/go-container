package slices

import (
	"github.com/shogo82148/go-container/tuples"
)

// Zip returns a slice of 2-tuples.
// The returned slice have the length of the shortest slice.
func Zip[S ~[]tuples.Tuple2[T1, T2], T1, T2 any](s1 []T1, s2 []T2) S {
	l := len(s1)
	if len(s2) < l {
		l = len(s2)
	}
	ret := make(S, l)
	for i := 0; i < l; i++ {
		ret[i] = tuples.Tuple2[T1, T2]{V1: s1[i], V2: s2[i]}
	}
	return ret
}
