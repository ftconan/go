// Author: magician
// File:   demo1.go
// Date:   2020/12/31
package main

const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

const (
	_        = iota
	KB int64 = 1 << (10 * iota)
	MB
	GB
	TB
)
const (
	A = iota
	B
	C = "c"
	D
	E = iota
	F
)

type Color int

const (
	Black Color = iota
	Red
	Blue
)

func test(c Color) {
	println(c)
}

func main() {
	c := Black
	test(c)

	//x := 1
	//test(x)

	test(1)
}
