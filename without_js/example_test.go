//go:build wasm && js

package without_js_test

import (
	"database/sql/driver"
	"fmt"
)

type A struct {
	Field string
}

func ExampleTest_one() {
	driver.IsValue(nil)
	fmt.Println(A{Field: "test1"})
	// Output:
	// {test1}
}

func ExampleTest_two() {
	fmt.Println(A{Field: "test2"})
	// Output:
	// {test2}
}
