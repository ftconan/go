// Author: magician
// File:   main.go
// Date:   2020/5/6
package main

import (
	"flag"
	"fmt"
	"strings"
)

var (
	n   = flag.Bool("n", false, "omit trailing newline")
	sep = flag.String("s", " ", "separator")
)

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}
