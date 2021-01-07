// Author: magician
// File:   demo1.go
// Date:   2021/1/3
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func NewTest() chan int {
	c := make(chan int)
	rand.Seed(time.Now().UnixNano())

	go func() {
		time.Sleep(time.Second)
		c <- rand.Int()
	}()

	return c
}

type Request struct {
	data []int
	ret  chan int
}

func NewRequest(data ...int) *Request {
	return &Request{data, make(chan int, 1)}
}

func Process(req *Request) {
	x := 0
	for _, i := range req.data {
		x += i
	}
	req.ret <- x
}

func main() {
	data := make(chan int)
	exit := make(chan bool)

	go func() {
		var i int
		for d := range data {
			fmt.Println(d)
			i++
		}
		fmt.Println("recv over.")
		exit <- true
	}()

	data <- 1
	data <- 2
	data <- 3
	close(data)
	fmt.Println("send over.")
	// TODO: WaitGroup
	time.Sleep(time.Second) // sleep -> recv over
	<-exit

	data1 := make(chan int, 3)
	exit1 := make(chan bool)
	data1 <- 1
	data1 <- 2
	data1 <- 3

	go func() {
		for d := range data1 {
			fmt.Println(d)
		}
		exit1 <- true
	}()

	data1 <- 4
	data1 <- 5
	close(data1)
	<-exit1

	d1 := make(chan int)
	d2 := make(chan int, 3)
	d2 <- 1
	fmt.Println(len(d1), cap(d1))
	fmt.Println(len(d2), cap(d2))

	// unidirectional
	c := make(chan int, 3)
	var send chan<- int = c // send-only
	var recv <-chan int = c // receive-only
	send <- 1
	//<-send
	<-recv
	//recv<-2

	// select
	//a, b := make(chan int, 3), make(chan int)
	//go func() {
	//	v, ok, s := 0, false, ""
	//
	//	for {
	//		select {
	//		case v, ok = <-a:
	//			s = "a"
	//		case v, ok = <-b:
	//			s = "b"
	//		}
	//		if ok {
	//			fmt.Println(s, v)
	//		} else {
	//			os.Exit(0)
	//		}
	//	}
	//}()
	//
	//for i := 0; i < 5; i++ {
	//	select {
	//	case a <- i:
	//	case b <- i:
	//	}
	//}
	//close(a)
	//select{} // 没有可⽤ channel，阻塞 main goroutine

	// mode
	t := NewTest()
	println(<-t)

	wg := sync.WaitGroup{}
	wg.Add(3)
	sem := make(chan int, 1)

	for i := 0; i < 3; i++ {
		go func(id int) {
			defer wg.Done()
			sem <- 1
			for x := 0; x < 3; x++ {
				fmt.Println(id, x)
			}
		}(i)
		<-sem
	}
	wg.Wait()

	// closed channel
	var wg1 sync.WaitGroup
	quit := make(chan bool)
	for i := 0; i < 2; i++ {
		wg1.Add(1)
		go func(id int) {
			defer wg1.Done()
			task := func() {
				println(id, time.Now().Nanosecond())
			}
			for {
				select {
				case <-quit:
					return
				default:
					task()
				}
			}
		}(i)
	}
	//time.Sleep(time.Second)
	close(quit)
	wg1.Wait()

	// select timeout
	w := make(chan bool)
	c1 := make(chan int, 2)
	go func() {
		select {
		case v := <-c1:
			fmt.Println(v)
		case <-time.After(time.Second * 3):
			fmt.Println("timeout.")
		}
		w <- true
	}()
	c1 <- 1
	<-w

	req := NewRequest(10, 20, 30)
	Process(req)
	fmt.Println(<-req.ret)
}
