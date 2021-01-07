// Author: magician
// File:   demo1.go
// Date:   2021/1/3
package main

import (
	"fmt"
	"runtime"
	"time"
	"unsafe"
)

type data struct {
	x [1024 * 100]byte
}

func test() uintptr {
	p := &data{}
	return uintptr(unsafe.Pointer(p))
}

func test1() unsafe.Pointer {
	p := &data{}
	return unsafe.Pointer(p)
}

type data1 struct {
	x [1024 * 100]byte
	y int
}

func test2() unsafe.Pointer {
	d := data1{}
	return unsafe.Pointer(&d.y)
}

type Data struct {
	d [1024 * 1000]byte
	o *Data
}

func test3() {
	var a, b Data
	a.o = &b
	b.o = &a
	runtime.SetFinalizer(&a, func(d *Data) { fmt.Printf("a %p final.\n", d) })
	runtime.SetFinalizer(&b, func(d *Data) { fmt.Printf("b %p final.\n", d) })
}

func main() {
	const N = 10000
	cache := new([N]uintptr)

	for i := 0; i < N; i++ {
		cache[i] = test()
		time.Sleep(time.Millisecond)
	}

	cache1 := new([N]unsafe.Pointer)

	for i := 0; i < N; i++ {
		//cache1[i] = test1()
		cache1[i] = test2()
		time.Sleep(time.Millisecond)
	}

	for {
		test3()
		time.Sleep(time.Millisecond)
	}
}
