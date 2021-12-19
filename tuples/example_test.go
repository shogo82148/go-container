package tuples_test

import (
	"fmt"

	"github.com/shogo82148/go-container/tuples"
)

func ExampleNew2() {
	t := tuples.New2(42, "foobar")
	fmt.Println(t.V1)
	fmt.Println(t.V2)

	//Output:
	// 42
	// foobar
}
func ExampleTuple2_Get() {
	t := tuples.New2(42, "foobar")
	v1, v2 := t.Get()
	fmt.Println(v1)
	fmt.Println(v2)

	//Output:
	// 42
	// foobar
}
