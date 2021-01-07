// Author: magician
// File:   demo1.go
// Date:   2021/1/2
package main

import "fmt"

type User struct {
	id int
	name string
}

type Manager struct {
	User
	title string
}

func (u *User) ToString() string  {
	return fmt.Sprintf("User: %p, %v", u, u)
}

func (m *Manager) ToString() string  {
	return fmt.Sprintf("Manager: %p, %v", m, m)
}

func main()  {
	m := Manager{User{1, "Tom"}, "Administrator"}
	//fmt.Printf("Manager: %p\n", &m)
	fmt.Println(m.ToString())
	fmt.Println(m.User.ToString())
}
