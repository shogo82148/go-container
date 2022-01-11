package list

import "github.com/shogo82148/go-container/tuples"

// Zip returns a list of 2-tuples.
// The returned lists have the length of the shortest list.
func Zip[T1, T2 any](l1 List[T1], l2 List[T2]) *List[tuples.Tuple2[T1, T2]] {
	ret := New[tuples.Tuple2[T1, T2]]()
	for e1, e2 := l1.Front(), l2.Front(); e1 != nil && e2 != nil; e1, e2 = e1.Next(), e2.Next() {
		ret.PushBack(tuples.Tuple2[T1, T2]{V1: e1.Value, V2: e2.Value})
	}
	return ret
}
