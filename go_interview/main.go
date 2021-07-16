// Author: magician
// File:   main.go
// Date:   2021/3/2
package main

import (
	"fmt"
	"sync"
)

type student struct {
	Name string
	Age  int
}

func pase_student() {
	m := make(map[string]*student)
	stus := [] student{
		{
			Name: "zhou",
			Age:  24,
		}, {
			Name: "li",
			Age:  23,
		}, {
			Name: "wang",
			Age:  22,
		},
	}
	// 错误写法
	//for _, stu := range stus {
	//	m[stu.Name] = &stu
	//}
	//for k, v := range m {
	//	println(k, "=>", v.Name)
	//}
	// 正确
	for i := 0; i < len(stus); i++ {
		m[stus[i].Name] = &stus[i]
	}
	for k, v := range m {
		println(k, "=>", v.Name)
	}
}

type People struct{}

func (p *People) ShowA() {
	fmt.Println("showA")
	p.ShowB()
}
func (p *People) ShowB() {
	fmt.Println("showB")
}

type Teacher struct {
	People
}

func (t *Teacher) ShowA() {
	fmt.Println("teacher ShowA")
	t.ShowB()
}

func (t *Teacher) ShowB() {
	fmt.Println("teacher showB")
}

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

type threadSafeSet struct {
	sync.RWMutex
	s []interface{}
}

func (set *threadSafeSet) Iter() <-chan interface{} {
	//ch := make(chan interface{})
	// 解除注释看看！
	ch := make(chan interface{}, len(set.s))
	go func() {
		set.RLock()
		for elem, value := range set.s {
			ch <- elem
			println("Iter:", elem, value)
		}
		close(ch)
		set.RUnlock()
	}()
	return ch
}

type People1 interface {
	Speak(string) string
}
type Student struct{}

func (stu *Student) Speak(think string) (talk string) {
	if think == "bitch" {
		talk = "You are a good boy"
	} else {
		talk = "hi"
	}
	return
}

func receiveMsg3(){
	msg := make(chan string)
	var wg sync.WaitGroup

	wg.Add(3)
	go func(){
		defer wg.Done()
		msg <- "goroutine 1"
	}()
	go func(){
		defer wg.Done()
		msg <- "goroutine 2"
	}()
	go func(){
		defer wg.Done()
		msg <- "goroutine 3"
	}()
	go func(){
		for i := range msg {
			fmt.Println("message :", i)
		}
	}()
	wg.Wait()
}

func main() {
	// for range(struct)
	//pase_student()

	// goroutine
	//runtime.GOMAXPROCS(1)
	//wg := sync.WaitGroup{}
	//wg.Add(20)
	//for i := 0; i < 10; i++ {
	//	go func() {
	//		fmt.Println("A: ", i)
	//		wg.Done()
	//	}()
	//}
	//for i := 0; i < 10; i++ {
	//	go func(i int) {
	//		fmt.Println("B: ", i)
	//		wg.Done()
	//	}(i)
	//}
	//wg.Wait()

	// struct
	//t := Teacher{}
	//t.ShowA()
	//t.ShowB()

	// defer
	//a := 1
	//b := 2
	//defer calc("1", a, calc("10", a, b))
	//a = 0
	//defer calc("2", a, calc("20", a, b))
	//b = 1

	// make, append
	//s := make([]int, 5)
	//s = append(s, 1, 2, 3)
	//fmt.Println(s)
	////
	//s1 := make([]int, 0)
	//s1 = append(s1, 1, 2, 3)
	//fmt.Println(s1)

	// sync, channel
	//th := threadSafeSet{
	//	s: []interface{}{
	//		"1", "2",
	//	},
	//}
	//v := <-th.Iter()
	//_ = fmt.Sprintf("%s%v", "ch", v)

	// *T
	//var peo People1 = Student{}
	//var peo People1 = &Student{}
	//think := "bitch"
	//fmt.Println(peo.Speak(think))
	//
	//total := 0
	//for i := 1; i <= 100; i++ {
	//	total += i
	//}
	//fmt.Println(total)

	//var m = map[int]string{
	//	1: "a",
	//	2: "b",
	//	3: "c",
	//}
	//for k, v := range m {
	//	fmt.Printf("%d:%v", k, v)
	//}

	//for i := 0; i < 3; i++ {
	//	defer fmt.Println(i)
	//}

	//receiveMsg3()
}
