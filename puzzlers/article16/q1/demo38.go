// Author: magician
// File:   demo38.go
// Date:   2021/1/15
package main

import "fmt"

func main()  {
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(i)
		}()
	}
}
