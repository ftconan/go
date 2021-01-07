// Author: magician
// File:   demo1.go
// Date:   2021/1/1
package main

import "fmt"

func main()  {
	m := map[int]struct{
		name string
		age int
	}{
		1: {"user1", 20},
		2: {"user2", 20},
	}
	println(m[1].name)

	//m1 := make(map[string]int, 1000)
	m1 := map[string]int{
		"a": 1,
	}
	if v, ok := m1["a"]; ok {
		println(v)
	}
	println(m1["c"])
	m1["b"] = 2
	delete(m1, "c")
	println(len(m1))
	for k, v := range m1 {
		println(k, v)
	}

	type user struct {
		name string
	}
	m3 := map[int]user{
		1: {"user1"},
	}
	//m3[1].name = "Tom"
	u := m3[1]
	u.name = "Tom"
	m3[1] = u
	fmt.Println(m3)

	m4 := map[int]*user{
		1: &user{"user1"},
	}
	m4[1].name = "Jack"
	fmt.Println(m4)

	for i := 0; i < 5; i++ {
		m := map[int]string{
			0: "a", 1: "a", 2: "a", 3: "a", 4: "a",
			5: "a", 6: "a", 7: "a", 8: "a", 9: "a",
		}
		for k := range m {
			m[k+k] = "x"
			delete(m, k)
		}
		fmt.Println(m)
	}
}
