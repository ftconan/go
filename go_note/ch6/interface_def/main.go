// Author: magician
// File:   demo1.go
// Date:   2021/1/2
package main

import "fmt"

type Stringer interface {
	String() string
}

type Printer interface {
	Stringer
	Print()
}

type Tester struct {
	s interface {
		String() string
	}
}

type User struct {
	id   int
	name string
}

func (u *User) String() string {
	return fmt.Sprintf("user %d, %s", u.id, u.name)
}

func (u *User) Print() {
	fmt.Println(u.String())
}

func Print(v interface{}) {
	fmt.Printf("%T: %v\n", v, v)
}

func main() {
	var t Printer = &User{1, "Tom"}
	t.Print()

	Print(1)
	Print("Hello, World!")

	//t := Tester{&User{1, "Tom"}}
	//fmt.Println(t.s.String())
}
