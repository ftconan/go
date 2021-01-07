// Author: magician
// File:   demo1.go
// Date:   2020/12/31
package main

import "fmt"

func main() {
	//var a struct{x int `a`}
	//var b struct{x int `ab`}
	//b = a

	type bigint int64
	var x bigint = 100
	println(x)

	x1 := 1234
	var b bigint = bigint(x1)
	var b2 int64 = int64(b)
	println(b2)

	type myslice []int
	var s myslice = []int{1, 2, 3}
	var s2 []int = s
	fmt.Println(s2)
}
