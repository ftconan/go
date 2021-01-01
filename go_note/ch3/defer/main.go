// Author: magician
// File:   main.go
// Date:   2020/12/31
package main

import (
	"os"
	"sync"
	"testing"
)

func test() error {
	f, err := os.Create("/Users/magician/Project/golang/go_note/ch3/defer/test.txt")
	if err != nil {
		return err
	}

	defer f.Close()

	_, _ = f.WriteString("Hello, World!")
	return nil
}

func test1(x int) {
	defer println("a")
	defer println("b")

	defer func() {
		println(100 / x)
	}()

	defer println("c")
}

func test2() {
	x, y := 10, 20
	defer func(i int) {
		println("defer:", i, y)
	}(x)

	x += 100
	y += 100
	println("x =", x, "y =", y)
}

var lock sync.Mutex

func test3() {
	lock.Lock()
	lock.Unlock()
}

func testdefer() {
	lock.Lock()
	defer lock.Unlock()
}

func BenchmarkTest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		test3()
	}
}

func BenchmarkTestDefer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testdefer()
	}
}

func main() {
	_ = test()
	//test1(0)
	test2()
}
