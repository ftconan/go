// Author: magician
// File:   demo1.go
// Date:   2021/1/1
package main

import "fmt"

func main()  {
	data := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	slice := data[1:4:5]
	fmt.Println(slice, data[:6:8], data[5:], data[:3], data[:])

	data1 := [...]int{0, 1, 2, 3, 4, 5}
	s := data1[2:4]
	s[0] += 100
	s[1] += 200
	fmt.Println(s)
	fmt.Println(data1)

	s1 := []int{0, 1, 2, 3, 8: 100}
	fmt.Println(s1, len(s1), cap(s1))
	s2 := make([]int, 6, 8)
	fmt.Println(s2, len(s2), cap(s2))
	s3 := make([]int, 6)
	fmt.Println(s3, len(s3), cap(s3))

	s4 := []int{0, 1, 2, 3}
	p := &s4[2]
	*p += 100
	fmt.Println(s4)

	d := [5]struct{
		x int
	}{}
	s5 := d[:]
	d[1].x = 10
	s5[2].x = 20
	fmt.Println(d)
	fmt.Printf("%p, %p\n", &d, &d[0])

	//
	s6 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s11 := s6[2:5]
	s22 := s11[2:6:7]
	//s33 := s22[3:6]
	fmt.Println(s11, s22)

	// reslice
	s111 := s6[2:5]
	s111[2] = 100
	s222 := s111[2:6]
	s222[3] = 200
	fmt.Println(s6)

	//append
	s7 := make([]int, 0, 5)
	fmt.Printf("%p\n", &s7)
	s2222 := append(s7, 1)
	fmt.Printf("%p\n", &s2222)

	s0 :=s6[:3]
	s22222 := append(s0, 100, 200)
	fmt.Println(s6)
	fmt.Println(s0)
	fmt.Println(s22222)

	data2 := [...]int{0, 1, 2, 3, 4, 10: 0}
	s8 := data2[:2:3]
	s8 = append(s8, 100, 200)
	fmt.Println(s8, data2)
	fmt.Println(&s8[0], &data2[0])

	s9 := make([]int, 0, 1)
	c := cap(s9)

	for i := 0; i < 50; i++ {
		s9 = append(s9, i)
		if n := cap(s9); n > c {
			fmt.Printf("cap: %d -> %d\n", c, n)
			c = n
		}
	}

	// copy
	data3 := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s10 := data3[8:]
	s12 := data3[:5]
	copy(s12, s10)
	fmt.Println(s12)
	fmt.Println(data3)
}
