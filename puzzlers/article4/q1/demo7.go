// Author: magician
// File:   demo7.go
// Date:   2021/1/9
package main

import (
	"flag"
	"fmt"
)

func main()  {
	//var name string
	//flag.StringVar(&name, "name", "everyone", "The greeting object.")

	var name = flag.String("name", "everyone", "The greeting object.")
	//name := flag.String("name", "everyone", "The greeting object.")


	flag.Parse()
	//fmt.Printf("Hello, %v!\n", name)
	fmt.Printf("Hello, %v!\n", *name)
}
