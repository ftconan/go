// Author: magician
// File:   demo10.go
// Date:   2021/1/9
package main

import "fmt"

//var block = "package"

func main()  {
	block := "function"
	{
		block := "inner"
		fmt.Printf("The block is %s!\n", block)
	}
	fmt.Printf("The block is %s!\n", block)
}
