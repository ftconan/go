// Author: magician
// File:   demo1.go
// Date:   2020/12/31
package main

func test() (int, int)  {
	return 1, 2
}

func add(x, y int) int  {
	return x + y
}

func sum(n ...int) int  {
	var x int
	for _, i := range n {
		x += i
	}
	return x
}

func add1(x, y int) (z int)  {
	z = x + y
	return
}

func add2(x, y int) (z int)  {
	var z1 = x + y
	return z1
}

func add3(x, y int) (z int)  {
	defer func() {
		z += 100
	}()

	z = x + y
	return
}

func add4(x, y int) (z int)  {
	defer func() {
		println(z)
	}()

	z = x + y
	return z + 200 // (z = z + 200) -> (call defer) -> (ret)
}

func main()  {
	//s := make([]int, 2)
	//s = test()

	x, _ := test()
	println(x)

	println(add(test()))
	println(sum(test()))

	// add
	println(add1(1, 2))
	println(add2(1, 2))
	println(add3(1, 2))
	println(add4(1, 2))
}
