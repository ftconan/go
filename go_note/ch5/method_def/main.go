// Author: magician
// File:   main.go
// Date:   2021/1/2
package main

import "fmt"

type Queue struct {
	elements []interface{}
}

func NewQueue() *Queue {
	return &Queue{make([]interface{}, 10)}
}

func (*Queue) Push(e interface{}) error {
	panic("not implemented")
}

//func (Queue) Push(e int) error  {
//	panic("not implemented")
//}

func (q *Queue) length() int {
	return len(q.elements)
}

type Data struct {
	x int
}

func (d Data) ValueTest() {
	fmt.Printf("Value: %p\n", &d)
}

func (d *Data) PointerTest() {
	fmt.Printf("Pointer: %p\n", d)
}

type X struct {
}

func (*X) test() {
	println("X.test")
}

func main() {
	d := Data{}
	p := &d
	fmt.Printf("Data: %p\n", p)
	d.ValueTest()
	d.PointerTest()
	p.ValueTest()
	p.PointerTest()

	p1 := &X{}
	p1.test()
	//(&p1).test()
}
