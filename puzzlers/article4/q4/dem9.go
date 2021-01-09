// Author: magician
// File:   dem9.go
// Date:   2021/1/9
package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	var err error
	n, err := io.WriteString(os.Stdout, "Hello, everyone!\n")
	if err != nil {
		fmt.Printf("Error %v\n", err)
	}
	fmt.Printf("%d byte(s) were written.\n", n)
}
