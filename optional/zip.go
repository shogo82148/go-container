package optional

import (
	"github.com/shogo82148/go-container/tuples"
)

// Zip returns an optional of 2-tuples.
func Zip[T optional[tuples.Tuple2[T1, T2]], T1, T2 any](v1 Optional[T1], v2 Optional[T2]) T {
	if v1.valid && v2.valid {
		return T{
			value: tuples.Tuple2[T1, T2]{v1.value, v2.value},
			valid: true,
		}
	}
	return T{}
}
