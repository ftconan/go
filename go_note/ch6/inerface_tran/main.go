// Author: magician
// File:   main.go
// Date:   2021/1/2
package main

import "fmt"

type Stringer interface {
	String() string
}

type Printer interface {
	String() string
	Print()
}

type User struct {
	id   int
	name string
}

func (u *User) String() string {
	return fmt.Sprintf("%d, %s", u.id, u.name)
}

func (u *User) Print() {
	fmt.Println(u.String())
}

type Tester interface {
	Do()
}

type FuncDo func()

func (d FuncDo) Do() {
	d()
}

func main() {
	var o interface{} = &User{1, "Tom"}
	if i, ok := o.(fmt.Stringer); ok {
		fmt.Println(i)
	}

	u := o.(*User)
	//u := o.(User)
	fmt.Println(u)

	switch v := o.(type) {
	case nil:
		fmt.Println("nil")
	case fmt.Stringer:
		fmt.Println(v)
	case func() string:
		fmt.Println(v)
	case *User:
		fmt.Printf("%d, %s\n", v.id, v.name)
	default:
		fmt.Println("unknown")
	}

	var o1 Printer = &User{1, "Tom"}
	var s Stringer = o1
	fmt.Println(s.String())

	// method_trick
	var t Tester = FuncDo(func() { println("Hello, World!") })
	t.Do()
}
