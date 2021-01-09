// Author: magician
// File:   demo6.go
// Date:   2021/1/9
package main

import (
	"flag"
	"golang/puzzlers/article3/q4/lib"
	//in "golang/puzzlers/article3/q2/lib/internal"
)

var name string

func init()  {
	flag.StringVar(&name, "name", "everyone", "The string object.")
}

func main()  {
	flag.Parse()
	lib.Hello(name)
}
