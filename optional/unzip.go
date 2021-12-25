package optional

import (
	"github.com/shogo82148/go-container/tuples"
)

// Unzip is an alias of Unzip2.
// It returns an optional of 2-tuples.
func Unzip[O1 optional[T1], O2 optional[T2], T1, T2 any](v Optional[tuples.Tuple2[T1, T2]]) (O1, O2) {
	if v.valid {
		return O1{
				value: v.value.V1,
				valid: true,
			}, O2{
				value: v.value.V2,
				valid: true,
			}
	}
	return O1{}, O2{}
}
