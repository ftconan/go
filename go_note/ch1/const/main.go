// Author: magician
// File:   demo1.go
// Date:   2020/12/31
package main

const x, y int = 1, 2
const s = "Hello, World!"
const (
	a, b = 10, 100
	c bool = false
)
//const (
//	s = "abc"
//	x
//)
//const (
//	a = "abc"
//	b = len(a)
//	c = unsafe.Sizeof(a)
//)
//const (
//	a byte = 100
//	b int = 1e20
//)

func main()  {
	//const x = "xxx"
	println(x, y)
	println(s)
	println(a, b, c)
}
