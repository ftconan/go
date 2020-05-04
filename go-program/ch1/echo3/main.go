// Author: magician
// File:   main.go
// Date:   2020/5/4
package main

import (
	"fmt"
	"os"
	"strings"
)

func main()  {
	fmt.Println(strings.Join(os.Args[1:], " "))
}
