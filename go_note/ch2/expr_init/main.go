// Author: magician
// File:   main.go
// Date:   2020/12/31
package main

import "fmt"

func length(s string) int {
	println("call length.")
	return len(s)
}

func main() {
	var a = struct {
		x int
	}{100}
	var b = []int{1, 2, 3}
	fmt.Println(a, b)

	a1 := []int{
		1,
		2,
	}
	b1 := []int{
		1,
		2}
	fmt.Println(a1, b1)

	// if
	x := 100
	if n := "abc"; x > 0 {
		println(n[2])
	} else if x < 0 {
		println(n[1])
	} else {
		println(n[0])
	}

	// for
	s := "abc"
	for i, n := 0, len(s); i < n; i++ {
		println(s[i])
	}

	n := len(s)
	for n > 0 {
		n--
		println(s[n])
	}

	s1 := "abcd"
	for i, n := 0, length(s1); i < n; i++ {
		println(i, s1[i])
	}

	// range
	s2 := "abc"
	for i := range s2 {
		println(s[i])
	}
	for _, c := range s2 {
		println(c)
	}
	for range s {
		println(s)
	}

	m := map[string]int{"a": 1, "b": 2}
	for k, v := range m {
		println(k, v)
	}

	a2 := [3]int{0, 1, 2}
	for i, v := range a2 {
		if i == 0 {
			a2[1], a2[2] = 999, 999
			fmt.Println(a2)
		}
		a2[i] = v + 100
	}
	fmt.Println(a2)

	s3 := []int{1, 2, 3, 4, 5}
	for i, v := range s3 {
		if i == 0 {
			s = s[:3]
			s3[2] = 100
		}
		println(i, v)
	}

	// switch
	x1 := []int{1, 2, 3}
	i := 2
	switch i {
	case x1[1]:
		println("a")
	case 1, 3:
		println("b")
	default:
		println("c")
	}

	// fallthrough
	x2 := 10
	switch x2 {
	case 10:
		println("a")
	case 0:
		println("b")
	}

	switch {
	case x1[1] > 0:
		println("a")
	case x1[1] < 0:
		println("b")
	default:
		println("c")
	}

	switch i := x1[2]; {
	case i > 0:
		println("a")
	case i < 0:
		println("b")
	default:
		println("c")
	}

	// goto, break, continue
	var i1 int
	for {
		println(i1)
		i1++
		if i1 > 2 {
			goto BREAK
		}
	}

BREAK:
	println("break")
	//EXIT:
	//	println("")

L1:
	for x := 0; x < 3; x++ {
	L2:
		for y := 0; y < 5; y++ {
			if y > 2 {
				continue L2
			}
			if x > 1 {
				break L1
			}

			print(x, ":", y, " ")
		}
		println()
	}

	x3 := 100
	switch {
	case x3 >= 0:
		if x3 == 0 {
			break
		}
		println(x3)
	}
}
