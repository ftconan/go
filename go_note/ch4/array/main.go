// Author: magician
// File:   main.go
// Date:   2021/1/1
package main

import "fmt"

func test(x [2]int) {
	fmt.Printf("x: %p\n", &x)
	x[1] = 100
}

func main() {
	a := [3]int{1, 2}
	b := [...]int{1, 2, 3, 4}
	c := [5]int{2: 100, 4: 200}
	d := [...]struct {
		name string
		age  uint8
	}{
		{"user1", 10},
		{"user2", 20},
	}
	fmt.Println(a, b, c, d)

	a1 := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	b1 := [...][2]int{{1, 1}, {2, 2}, {3, 3}}
	fmt.Println(a1, b1)

	// test
	a2 := [2]int{}
	fmt.Printf("a: %p\n", &a2)
	test(a2)
	fmt.Println(a2)

	// len, cap
	println(len(a2), cap(a2))
}
