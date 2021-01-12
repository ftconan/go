// Author: magician
// File:   demo20.go
// Date:   2021/1/12
package main

import "fmt"

func main() {
	ch1 := make(chan int, 3)
	ch1 <- 2
	ch1 <- 1
	ch1 <- 3
	elem1 := <-ch1
	fmt.Printf("The first element received from channel ch1: %v\n", elem1)

	//defer close(ch1)
	//go func() {
	//	for d := range ch1 {
	//		fmt.Println(d)
	//	}
	//}()
}
