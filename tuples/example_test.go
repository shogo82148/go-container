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
