// Author: magician
// File:   demo26.go
// Date:   2021/1/12
package main

import "fmt"

type Printer func(contents string) (n int, err error)

func printToStd(contents string) (byteNum int, err error)  {
	return fmt.Println(contents)
}

func main()  {
	var p Printer
	p = printToStd
	_, _ = p("something")
}
