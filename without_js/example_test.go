//go:build wasm && js

package without_js_test

import (
	"fmt"
)

type A struct {
	Field string
}

func ExampleTest_one() {
	fmt.Println(A{Field: "test1"})
	// Output:
	// {test1}
}

func ExampleTest_two() {
	fmt.Println(A{Field: "test2"})
	// Output:
	// {test2}
}
