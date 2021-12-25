package slices

import (
	"github.com/shogo82148/go-container/tuples"
)

// Unzip is an alias of Unzip2.
// It converts a slice of 2-tuples to slices of each elements.
func Unzip[S1 ~[]T1, S2 ~[]T2, T1, T2 any](s []tuples.Tuple2[T1, T2]) (S1, S2) {
	s1 := make(S1, len(s))
	s2 := make(S2, len(s))
	for i, t := range s {
		s1[i] = t.V1
		s2[i] = t.V2
	}
	return s1, s2
}
