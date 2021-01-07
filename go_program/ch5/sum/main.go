// Author: magician
// File:   demo1.go
// Date:   2020/6/17
package main

import "fmt"

func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

func main() {
	fmt.Println(sum())
	fmt.Println(sum(3))
	fmt.Println(sum(1, 2, 3, 4))

	// slice
	values := []int{1, 2, 3, 4}
	fmt.Println(sum(values...))
}
