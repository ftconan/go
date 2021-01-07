// Author: magician
// File:   demo1.go
// Date:   2020/12/31
package main

import "fmt"

func test() func()  {
	x := 100
	fmt.Printf("x (%p) = %d\n", &x, x)

	return func() {
		fmt.Printf("x (%p) = %d\n", &x, x)
	}
}

func main() {
	// --- function variable ---
	fn := func() {
		println("Hello, World!")
	}
	fn()

	// --- function collection ---
	fns := []func(x int) int{
		func(x int) int {
			return x + 1
		},
		func(x int) int {
			return x + 2
		},
	}
	println(fns[0](100))

	// --- function as field ---
	d := struct {
		fn func() string
	}{
		fn: func() string {
			return "Hello, World!"
		},
	}
	println(d.fn())

	// --- channel of function ---
	fc := make(chan func() string, 2)
	fc <- func() string {
		return "Hello, World!"
	}
	println((<-fc)())

	// test
	f := test()
	f()
}
