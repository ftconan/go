// Author: magician
// File:   demo5.go
// Date:   2021/1/9
package main

import (
	"flag"
	"golang/puzzlers/article3/q2/lib"
)

var name string

func init()  {
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
}

func main()  {
	flag.Parse()
	lib.Hello(name)
}
