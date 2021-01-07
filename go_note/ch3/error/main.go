// Author: magician
// File:   demo1.go
// Date:   2020/12/31
package main

import (
	"errors"
	"fmt"
)

func test() {
	defer func() {
		if err := recover(); err != nil {
			println(err.(string))
		}
	}()

	panic("panic error!")
}

func test1() {
	defer func() {
		fmt.Println(recover())
	}()

	defer func() {
		panic("defer panic")
	}()

	panic("test panic")
}

func test2() {
	//defer recover()
	defer fmt.Println(recover())
	defer func() {
		func() {
			println("defer inner")
			recover()
		}()
	}()
	//panic("test panic")
}

func expect() {
	recover()
}

func test3() {
	defer expect()
	panic("test panic")
}

func test4(x, y int) {
	var z int
	func() {
		defer func() {
			if recover() != nil {
				z = 0
			}
		}()

		z = x / y
		return
	}()

	println("x / y =", z)
}

var ErrDivByZero = errors.New("division by zero")

func div(x, y int) (int, error) {
	if y == 0 {
		return 0, ErrDivByZero
	}
	return x / y, nil
}

func main() {
	test()
	test1()
	test2()
	test3()
	test4(1, 0)

	// div
	switch z, err := div(10, 0); err {
	case nil:
		println(z)
	case ErrDivByZero:
		panic(err)
	}
}
