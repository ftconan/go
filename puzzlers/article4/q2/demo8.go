// Author: magician
// File:   demo8.go
// Date:   2021/1/9
package main

import (
	"flag"
	"fmt"
)

func main()  {
	var name = getTheFlag()
	flag.Parse()
	fmt.Printf("Hello, %v\n", *name)
}

func getTheFlag() *string {
	return flag.String("name", "everyone", "The greeting object.")
}

//func getTheFlag() *int {
//	return flag.Int("name", 1, "The greeting object.")
//}
