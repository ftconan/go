// Author: magician
// File:   demo2.go
// Date:   2021/1/7
package main

import (
	"flag"
	"fmt"
)

var name string

func init()  {
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
}

func main()  {
	flag.Parse()
	fmt.Printf("Hello %s!\n", name)
}
