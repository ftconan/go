// Author: magician
// File:   tempflag.go
// Date:   2020/7/8
package main

import (
	"flag"
	"fmt"
	"golang/go_program/ch7/tempconv"
)

var temp = tempconv.CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}
