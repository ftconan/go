// Author: magician
// File:   main.go
// Date:   2021/1/2
package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type User struct {
	id   int
	name string
}

func main() {
	u := User{1, "Tom"}
	var i interface{} = u
	u.id = 2
	u.name = "Jack"
	fmt.Printf("%v\n", u)
	fmt.Printf("%v\n", i.(User))

	u1 := User{1, "Tom"}
	var vi, pi interface{} = u1, &u1
	//vi.(User).name = "Jack"
	pi.(*User).name = "Jack"
	fmt.Printf("%v\n", vi.(User))
	fmt.Printf("%v\n", pi.(*User))

	var a interface{} = nil
	var b interface{} = (*int)(nil)

	type iface struct {
		itab, data uintptr
	}

	ia := (*iface)(unsafe.Pointer(&a))
	ib := (*iface)(unsafe.Pointer(&b))
	fmt.Println(a == nil, ia)
	fmt.Println(b == nil, ib, reflect.ValueOf(b).IsNil())
}
