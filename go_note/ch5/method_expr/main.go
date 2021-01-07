// Author: magician
// File:   demo1.go
// Date:   2021/1/2
package main

import "fmt"

type User struct {
	id   int
	name string
}

func (u *User) Test() {
	fmt.Printf("%p, %v\n", u, u)
}

func (u User) Test1() {
	fmt.Println(u)
}

func (u *User) TestPointer() {
	fmt.Printf("TestPointer: %p, %v\n", u, u)
}

func (u User) TestValue() {
	fmt.Printf("TestValue: %p, %v\n", &u, u)
}

type Data struct {
}

func (Data) TestValue() {

}

func (*Data) TestPointer() {

}

func main() {
	u := User{1, "Tom"}
	u.Test()

	mValue := u.Test
	mValue()

	mExpression := (*User).Test
	mExpression(&u)

	// T method -> copy value
	u1 := User{1, "Tom"}
	mValue1 := u1.Test1
	u1.id, u1.name = 2, "Jack"
	u1.Test1()
	mValue1()

	fmt.Printf("User: %p, %v\n", &u, u)
	mv := User.TestValue
	mv(u)
	mp := (*User).TestPointer
	mp(&u)
	mp2 := (*User).TestValue
	mp2(&u)

	var p *Data = nil
	p.TestPointer()
	(*Data)(nil).TestPointer() //method value
	(*Data).TestPointer(nil)   // method expression
	//p.TestValue()
	//(Data)(nil).TestValue()
	//Data.TestValue(nil)
}
