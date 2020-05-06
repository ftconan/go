// Author: magician
// File:   main.go
// Date:   2020/5/6
package main

import "fmt"

func main() {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%g°F = %g°C\n", freezingF, fToC(freezingF)) // "32°F = 0°C"
	fmt.Printf("%g°F = %g°C\n", freezingF, fToC(boilingF))  // "212°F = 100°C"
}

func fToC(f float64) interface{} {
	return (f - 32) * 5 / 9
}
