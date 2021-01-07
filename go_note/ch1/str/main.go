// Author: magician
// File:   demo1.go
// Date:   2020/12/31
package main

import "fmt"

func main() {
	s := "abc"
	println(s[0] == '\x61', s[1] == 'b', s[2] == 0x63)

	s1 := `a
	b\r\n\x00
	c`
	println(s1)

	s2 := "Hello, " +
		"World!"
	println(s2)

	s3 := "Hello, World!"
	s31 := s3[:5]
	s32 := s3[7:]
	s33 := s3[1:5]
	println(s3, s31, s32, s33)

	fmt.Printf("%T\n", 'a')
	var c1, c2 rune = '\u6211', '们'
	println(c1 == '我', string(c2) == "\xe4\xbb\xac")

	s4 := "abcd"
	bs := []byte(s)
	bs[1] = 'B'
	println(s4, string(bs))

	u := "电脑"
	us := []rune(u)
	us[1] = '话'
	println(string(us))

	s5 := "abc汉字"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c,", s5[i]) // byte
	}
	fmt.Println()
	for _, r := range s5 {
		fmt.Printf("%c", r) // rune
	}

}
