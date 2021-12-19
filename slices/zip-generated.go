// Code generated by generate-zip.pl; DO NOT EDIT.

package slices

import (
	"fmt"

	"github.com/shogo82148/go-container/tuples"
)

// Zip2 returns a slice of 2-tuples.
// All slices must have same length. the lengths are different, Zip2 panics
func Zip2[T1, T2 any](s1 []T1, s2 []T2) []tuples.Tuple2[T1, T2] {
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

// Zip3 returns a slice of 3-tuples.
// All slices must have same length. the lengths are different, Zip3 panics
func Zip3[T1, T2, T3 any](s1 []T1, s2 []T2, s3 []T3) []tuples.Tuple3[T1, T2, T3] {
	l := len(s1)
	if len(s2) != l {
		panic(fmt.Errorf("s2 have different length from s1. len(s2) = %d, len(s1) = %d", len(s2), l))
	}
	if len(s3) != l {
		panic(fmt.Errorf("s3 have different length from s1. len(s3) = %d, len(s1) = %d", len(s3), l))
	}
	ret := make([]tuples.Tuple3[T1, T2, T3], l)
	for i := 0; i < l; i++ {
		ret[i] = tuples.Tuple3[T1, T2, T3]{s1[i], s2[i], s3[i]}
	}
	return ret
}

// Zip4 returns a slice of 4-tuples.
// All slices must have same length. the lengths are different, Zip4 panics
func Zip4[T1, T2, T3, T4 any](s1 []T1, s2 []T2, s3 []T3, s4 []T4) []tuples.Tuple4[T1, T2, T3, T4] {
	l := len(s1)
	if len(s2) != l {
		panic(fmt.Errorf("s2 have different length from s1. len(s2) = %d, len(s1) = %d", len(s2), l))
	}
	if len(s3) != l {
		panic(fmt.Errorf("s3 have different length from s1. len(s3) = %d, len(s1) = %d", len(s3), l))
	}
	if len(s4) != l {
		panic(fmt.Errorf("s4 have different length from s1. len(s4) = %d, len(s1) = %d", len(s4), l))
	}
	ret := make([]tuples.Tuple4[T1, T2, T3, T4], l)
	for i := 0; i < l; i++ {
		ret[i] = tuples.Tuple4[T1, T2, T3, T4]{s1[i], s2[i], s3[i], s4[i]}
	}
	return ret
}

// Zip5 returns a slice of 5-tuples.
// All slices must have same length. the lengths are different, Zip5 panics
func Zip5[T1, T2, T3, T4, T5 any](s1 []T1, s2 []T2, s3 []T3, s4 []T4, s5 []T5) []tuples.Tuple5[T1, T2, T3, T4, T5] {
	l := len(s1)
	if len(s2) != l {
		panic(fmt.Errorf("s2 have different length from s1. len(s2) = %d, len(s1) = %d", len(s2), l))
	}
	if len(s3) != l {
		panic(fmt.Errorf("s3 have different length from s1. len(s3) = %d, len(s1) = %d", len(s3), l))
	}
	if len(s4) != l {
		panic(fmt.Errorf("s4 have different length from s1. len(s4) = %d, len(s1) = %d", len(s4), l))
	}
	if len(s5) != l {
		panic(fmt.Errorf("s5 have different length from s1. len(s5) = %d, len(s1) = %d", len(s5), l))
	}
	ret := make([]tuples.Tuple5[T1, T2, T3, T4, T5], l)
	for i := 0; i < l; i++ {
		ret[i] = tuples.Tuple5[T1, T2, T3, T4, T5]{s1[i], s2[i], s3[i], s4[i], s5[i]}
	}
	return ret
}

// Zip6 returns a slice of 6-tuples.
// All slices must have same length. the lengths are different, Zip6 panics
func Zip6[T1, T2, T3, T4, T5, T6 any](s1 []T1, s2 []T2, s3 []T3, s4 []T4, s5 []T5, s6 []T6) []tuples.Tuple6[T1, T2, T3, T4, T5, T6] {
	l := len(s1)
	if len(s2) != l {
		panic(fmt.Errorf("s2 have different length from s1. len(s2) = %d, len(s1) = %d", len(s2), l))
	}
	if len(s3) != l {
		panic(fmt.Errorf("s3 have different length from s1. len(s3) = %d, len(s1) = %d", len(s3), l))
	}
	if len(s4) != l {
		panic(fmt.Errorf("s4 have different length from s1. len(s4) = %d, len(s1) = %d", len(s4), l))
	}
	if len(s5) != l {
		panic(fmt.Errorf("s5 have different length from s1. len(s5) = %d, len(s1) = %d", len(s5), l))
	}
	if len(s6) != l {
		panic(fmt.Errorf("s6 have different length from s1. len(s6) = %d, len(s1) = %d", len(s6), l))
	}
	ret := make([]tuples.Tuple6[T1, T2, T3, T4, T5, T6], l)
	for i := 0; i < l; i++ {
		ret[i] = tuples.Tuple6[T1, T2, T3, T4, T5, T6]{s1[i], s2[i], s3[i], s4[i], s5[i], s6[i]}
	}
	return ret
}

// Zip7 returns a slice of 7-tuples.
// All slices must have same length. the lengths are different, Zip7 panics
func Zip7[T1, T2, T3, T4, T5, T6, T7 any](s1 []T1, s2 []T2, s3 []T3, s4 []T4, s5 []T5, s6 []T6, s7 []T7) []tuples.Tuple7[T1, T2, T3, T4, T5, T6, T7] {
	l := len(s1)
	if len(s2) != l {
		panic(fmt.Errorf("s2 have different length from s1. len(s2) = %d, len(s1) = %d", len(s2), l))
	}
	if len(s3) != l {
		panic(fmt.Errorf("s3 have different length from s1. len(s3) = %d, len(s1) = %d", len(s3), l))
	}
	if len(s4) != l {
		panic(fmt.Errorf("s4 have different length from s1. len(s4) = %d, len(s1) = %d", len(s4), l))
	}
	if len(s5) != l {
		panic(fmt.Errorf("s5 have different length from s1. len(s5) = %d, len(s1) = %d", len(s5), l))
	}
	if len(s6) != l {
		panic(fmt.Errorf("s6 have different length from s1. len(s6) = %d, len(s1) = %d", len(s6), l))
	}
	if len(s7) != l {
		panic(fmt.Errorf("s7 have different length from s1. len(s7) = %d, len(s1) = %d", len(s7), l))
	}
	ret := make([]tuples.Tuple7[T1, T2, T3, T4, T5, T6, T7], l)
	for i := 0; i < l; i++ {
		ret[i] = tuples.Tuple7[T1, T2, T3, T4, T5, T6, T7]{s1[i], s2[i], s3[i], s4[i], s5[i], s6[i], s7[i]}
	}
	return ret
}

// Zip8 returns a slice of 8-tuples.
// All slices must have same length. the lengths are different, Zip8 panics
func Zip8[T1, T2, T3, T4, T5, T6, T7, T8 any](s1 []T1, s2 []T2, s3 []T3, s4 []T4, s5 []T5, s6 []T6, s7 []T7, s8 []T8) []tuples.Tuple8[T1, T2, T3, T4, T5, T6, T7, T8] {
	l := len(s1)
	if len(s2) != l {
		panic(fmt.Errorf("s2 have different length from s1. len(s2) = %d, len(s1) = %d", len(s2), l))
	}
	if len(s3) != l {
		panic(fmt.Errorf("s3 have different length from s1. len(s3) = %d, len(s1) = %d", len(s3), l))
	}
	if len(s4) != l {
		panic(fmt.Errorf("s4 have different length from s1. len(s4) = %d, len(s1) = %d", len(s4), l))
	}
	if len(s5) != l {
		panic(fmt.Errorf("s5 have different length from s1. len(s5) = %d, len(s1) = %d", len(s5), l))
	}
	if len(s6) != l {
		panic(fmt.Errorf("s6 have different length from s1. len(s6) = %d, len(s1) = %d", len(s6), l))
	}
	if len(s7) != l {
		panic(fmt.Errorf("s7 have different length from s1. len(s7) = %d, len(s1) = %d", len(s7), l))
	}
	if len(s8) != l {
		panic(fmt.Errorf("s8 have different length from s1. len(s8) = %d, len(s1) = %d", len(s8), l))
	}
	ret := make([]tuples.Tuple8[T1, T2, T3, T4, T5, T6, T7, T8], l)
	for i := 0; i < l; i++ {
		ret[i] = tuples.Tuple8[T1, T2, T3, T4, T5, T6, T7, T8]{s1[i], s2[i], s3[i], s4[i], s5[i], s6[i], s7[i], s8[i]}
	}
	return ret
}

// Zip9 returns a slice of 9-tuples.
// All slices must have same length. the lengths are different, Zip9 panics
func Zip9[T1, T2, T3, T4, T5, T6, T7, T8, T9 any](s1 []T1, s2 []T2, s3 []T3, s4 []T4, s5 []T5, s6 []T6, s7 []T7, s8 []T8, s9 []T9) []tuples.Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9] {
	l := len(s1)
	if len(s2) != l {
		panic(fmt.Errorf("s2 have different length from s1. len(s2) = %d, len(s1) = %d", len(s2), l))
	}
	if len(s3) != l {
		panic(fmt.Errorf("s3 have different length from s1. len(s3) = %d, len(s1) = %d", len(s3), l))
	}
	if len(s4) != l {
		panic(fmt.Errorf("s4 have different length from s1. len(s4) = %d, len(s1) = %d", len(s4), l))
	}
	if len(s5) != l {
		panic(fmt.Errorf("s5 have different length from s1. len(s5) = %d, len(s1) = %d", len(s5), l))
	}
	if len(s6) != l {
		panic(fmt.Errorf("s6 have different length from s1. len(s6) = %d, len(s1) = %d", len(s6), l))
	}
	if len(s7) != l {
		panic(fmt.Errorf("s7 have different length from s1. len(s7) = %d, len(s1) = %d", len(s7), l))
	}
	if len(s8) != l {
		panic(fmt.Errorf("s8 have different length from s1. len(s8) = %d, len(s1) = %d", len(s8), l))
	}
	if len(s9) != l {
		panic(fmt.Errorf("s9 have different length from s1. len(s9) = %d, len(s1) = %d", len(s9), l))
	}
	ret := make([]tuples.Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9], l)
	for i := 0; i < l; i++ {
		ret[i] = tuples.Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]{s1[i], s2[i], s3[i], s4[i], s5[i], s6[i], s7[i], s8[i], s9[i]}
	}
	return ret
}

// Zip10 returns a slice of 10-tuples.
// All slices must have same length. the lengths are different, Zip10 panics
func Zip10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10 any](s1 []T1, s2 []T2, s3 []T3, s4 []T4, s5 []T5, s6 []T6, s7 []T7, s8 []T8, s9 []T9, s10 []T10) []tuples.Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10] {
	l := len(s1)
	if len(s2) != l {
		panic(fmt.Errorf("s2 have different length from s1. len(s2) = %d, len(s1) = %d", len(s2), l))
	}
	if len(s3) != l {
		panic(fmt.Errorf("s3 have different length from s1. len(s3) = %d, len(s1) = %d", len(s3), l))
	}
	if len(s4) != l {
		panic(fmt.Errorf("s4 have different length from s1. len(s4) = %d, len(s1) = %d", len(s4), l))
	}
	if len(s5) != l {
		panic(fmt.Errorf("s5 have different length from s1. len(s5) = %d, len(s1) = %d", len(s5), l))
	}
	if len(s6) != l {
		panic(fmt.Errorf("s6 have different length from s1. len(s6) = %d, len(s1) = %d", len(s6), l))
	}
	if len(s7) != l {
		panic(fmt.Errorf("s7 have different length from s1. len(s7) = %d, len(s1) = %d", len(s7), l))
	}
	if len(s8) != l {
		panic(fmt.Errorf("s8 have different length from s1. len(s8) = %d, len(s1) = %d", len(s8), l))
	}
	if len(s9) != l {
		panic(fmt.Errorf("s9 have different length from s1. len(s9) = %d, len(s1) = %d", len(s9), l))
	}
	if len(s10) != l {
		panic(fmt.Errorf("s10 have different length from s1. len(s10) = %d, len(s1) = %d", len(s10), l))
	}
	ret := make([]tuples.Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10], l)
	for i := 0; i < l; i++ {
		ret[i] = tuples.Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]{s1[i], s2[i], s3[i], s4[i], s5[i], s6[i], s7[i], s8[i], s9[i], s10[i]}
	}
	return ret
}
