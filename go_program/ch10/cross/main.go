// Author: magician
// File:   main.go
// Date:   2020/7/30
package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.GOOS, runtime.GOARCH)
}
