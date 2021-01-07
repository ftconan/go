// Author: magician
// File:   demo1.go
// Date:   2020/12/31
package main

import (
	"fmt"
	"unsafe"
)

func test() *int {
	x := 100
	return &x
}

func main() {
	type data struct {
		a int
	}
	var d = data{1234}
	var p *data
	p = &d
	fmt.Printf("%p, %v\n", p, p.a)

	x := 1234
	p1 := &x
	fmt.Printf("%T\n", p1)
	//p1++

	x1 := 0x12345678
	p2 := unsafe.Pointer(&x1) // *int -> Pointer
	n := (*[4]byte)(p2)       // Pointer -> *[4]byte
	for i := 0; i < len(n); i++ {
		fmt.Printf("%X ", n[i])
	}

	fmt.Printf("%T\n", test())

	d1 := struct {
		s string
		x int
	}{"abc", 100}
	p3 := uintptr(unsafe.Pointer(&d)) // *struct -> Pointer -> uintptr
	p3 += unsafe.Offsetof(d1.x)       // uintptr + offset
	p4 := unsafe.Pointer(p3)          //unitptr -> Pointer
	px := (*int)(p4)                  // Pointer -> *int
	*px = 200                         // d1.x = 200
	fmt.Printf("%#v\n", d1)
}
