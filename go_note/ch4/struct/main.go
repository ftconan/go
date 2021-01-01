// Author: magician
// File:   main.go
// Date:   2021/1/1
package main

import "fmt"

type Node struct {
	_    int
	id   int
	data *byte
	next *Node
}

func main() {
	n1 := Node{
		id:   1,
		data: nil,
	}
	n2 := Node{
		id:   2,
		data: nil,
		next: &n1,
	}
	fmt.Println(n1, n2)

	type User struct {
		name string
		age  int
	}
	u1 := User{"Tom", 20}
	//u2 := User{"Tom"}
	fmt.Println(u1)

	type File struct {
		name string
		size int
		attr struct {
			perm  int
			owner int
		}
	}
	f := File{
		name: "test.txt",
		size: 1025,
		//attr: {0755, 1},
	}
	f.attr.owner = 1
	f.attr.perm = 0755
	fmt.Println(f)
	var attr = struct {
		perm  int
		owner int
	}{
		2,
		0775,
	}
	f.attr = attr
	fmt.Println(f)

	type User1 struct {
		id   int
		name string
	}
	m := map[User1]int{
		User1{1, "Tom"}: 100,
	}
	fmt.Println(m)

	//var u1 struct{ name string "username" }
	//var u2 struct{ name string }
	//u2 = u1

	// set
	var null struct{}
	set := make(map[string]struct{})
	set["a"] = null
	fmt.Println(set)

	// anonymous
	type User2 struct {
		name string
	}
	type Manager struct {
		User2
		title string
	}
	m1 := Manager{
		User2: User2{"Tom"},
		title: "Administrator",
	}
	fmt.Println(m1)

	type Resource struct {
		id int
	}
	type User3 struct {
		Resource
		name string
	}
	type Manager1 struct {
		User3
		title string
	}
	var m2 Manager1
	m2.id = 1
	m2.name = "Jack"
	m2.title = "Administrator"
	fmt.Println(m2)

	type Resource1 struct {
		id   int
		name string
	}
	type Classify struct {
		id int
	}
	type User4 struct {
		Resource1
		Classify
		name string
	}
	u := User4{
		Resource1{1, "people"},
		Classify{100},
		"Jack",
	}
	println(u.name)
	println(u.Resource1.name)
	//println(u.id)
	println(u.Resource1.id)

	type Resource2 struct {
		id int
	}
	type User5 struct {
		*Resource2
		//Resource2
		name string
	}
	u2 := User5{
		&Resource2{1},
		"Administrator",
	}
	fmt.Println(u2)

	// oop
	type User6 struct {
		id   int
		name string
	}
	type Manager3 struct {
		User6
		title string
	}
	m3 := Manager3{User6{1, "Tom"}, "Administrator"}
	//var u3 User = m3
	var u3 User6 = m3.User6
	fmt.Println(u3)
}
