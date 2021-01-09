// Author: magician
// File:   demo4.go
// Date:   2021/1/9
package main

import "flag"

var name string

func init()  {
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
}

func main()  {
	flag.Parse()
	hello(name)
}
