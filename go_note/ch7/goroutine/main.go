// Author: magician
// File:   main.go
// Date:   2021/1/3
package main

import (
	"fmt"
	"math"
	"runtime"
	"sync"
)

func sum(id int) {
	var x int64
	for i := 0; i < math.MaxUint32; i++ {
		x += int64(i)
	}
	println(id, x)
}

func main() {
	go func() {
		fmt.Println("Hello, World!")
	}()

	wg := new(sync.WaitGroup)
	wg.Add(2)

	for i := 0; i < 2; i++ {
		go func(id int) {
			defer wg.Done()
			sum(id)
		}(i)
	}
	wg.Wait()

	// runtime.Goexit
	wg1 := new(sync.WaitGroup)
	wg1.Add(1)

	go func() {
		defer wg1.Done()
		defer println("A.defer")

		func() {
			defer println("B.defer")
			runtime.Goexit() // exit
			println("B")
		}()
		println("A")
	}()

	wg1.Wait()

	// python: yield  -> runtime.Gosched
	wg2 := new(sync.WaitGroup)
	wg2.Add(2)

	go func() {
		defer wg2.Done()
		for i := 0; i < 6; i++ {
			println(i)
			if i == 3 {
				runtime.Gosched()
			}
		}
	}()

	go func() {
		defer wg2.Done()
		println("Hello, World!")
	}()

	wg2.Wait()
}
