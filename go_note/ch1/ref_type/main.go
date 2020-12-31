// Author: magician
// File:   main.go
// Date:   2020/12/31
package main

import "fmt"

func main()  {
	a := []int{0, 0, 0}
	a[1] = 10
	println(a)
	fmt.Println(a)

	b := make([]int, 3)
	b[1] = 10
	println(b)
	fmt.Println(b)

	c := new([]int)
	//c[1] = 10
	println(c)
	fmt.Println(c)
}
