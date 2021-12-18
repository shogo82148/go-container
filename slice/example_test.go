package slice_test

import (
	"fmt"

	"github.com/shogo82148/go-container/slice"
)

func ExampleMap() {
	out := slice.Map(
		[]int{1, 2, 3, 4, 5, 6},
		func(a int) string {
			return fmt.Sprintf("<%d>", a)
		},
	)

	for _, v := range out {
		fmt.Println(v)
	}

	//Output:
	// <1>
	// <2>
	// <3>
	// <4>
	// <5>
	// <6>
}
