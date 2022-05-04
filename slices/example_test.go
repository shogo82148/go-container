package slices_test

import (
	"fmt"

	"github.com/shogo82148/go-container/slices"
	"github.com/shogo82148/go-container/tuples"
)

func ExampleFilter() {
	out := slices.Filter(
		[]int{1, 2, 3, 4, 5, 6},
		func(a int) bool {
			return a > 3
		},
	)

	for _, v := range out {
		fmt.Println(v)
	}

	//Output:
	// 4
	// 5
	// 6
}

func ExampleMap() {
	out := slices.Map[[]string](
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

func ExampleZip() {
	out := slices.Zip[[]tuples.Tuple2[int, string]](
		[]int{1, 2, 3, 4, 5},
		[]string{"one", "two", "three", "four", "five"},
	)
	for _, tuple := range out {
		fmt.Printf("%d => %q\n", tuple.V1, tuple.V2)
	}

	//Output:
	// 1 => "one"
	// 2 => "two"
	// 3 => "three"
	// 4 => "four"
	// 5 => "five"
}
