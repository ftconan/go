// Author: magician
// File:   demo27.go
// Date:   2021/1/12
package main

import (
	"errors"
	"fmt"
)

type operator func(x, y int) int

func calculate(x, y int, op operator) (int, error) {
	if op == nil {
		return 0, errors.New("invalid operation")
	}
	return op(x, y), nil
}

type calculateFunc func(x int, y int) (int, error)

func getCalculator(op operator) calculateFunc  {
	return func(x, y int) (int, error) {
		if op == nil {
			return 0, errors.New("invalid operation")
		}
		return op(x, y), nil
	}
}

func main() {
	x, y := 12, 23
	op := func(x, y int) int {
		return x + y
	}
	result, err := calculate(x, y, op)
	fmt.Printf("The result: %d (error: %v)\n", result, err)
	result, err = calculate(x, y, nil)
	fmt.Printf("The result: %d (error: %v)\n", result, err)

	x, y = 56, 78
	add := getCalculator(op)
	result, err = add(x, y)
	fmt.Printf("The result: %d (error: %v)\n", result, err)
}
