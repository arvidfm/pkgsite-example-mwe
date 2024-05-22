//go:build wasm && js

package with_js_test

import (
	"fmt"
	"syscall/js"
)

type A struct {
	Field string
}

func ExampleTest_one() {
	js.ValueOf(nil)
	fmt.Println(A{Field: "test1"})
	// Output:
	// {test1}
}

func ExampleTest_two() {
	fmt.Println(A{Field: "test2"})
	// Output:
	// {test2}
}
