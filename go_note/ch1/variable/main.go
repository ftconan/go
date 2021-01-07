// Author: magician
// File:   demo1.go
// Date:   2020/12/30
package main

//var x int
//var f float32 = 6
//var s = "abc"

func test() (int, string) {
	return 1, "abc"
}

func main() {
	x := 123
	n, s := 0x1234, "Hello World!"
	println(x, s, n)

	data, i := [3]int{1, 2, 3}, 0
	i, data[i] = 2, 100

	_, s = test()
	println(s)

	s = "abc"
	println(&s)
	s, y := "hello", 20
	println(&s, y)

	{
		s, z := 1000, 30
		println(&s, z)
	}

}
