// Author: magician
// File:   demo13.go
// Date:   2021/1/10
package main

import "fmt"

func main() {
	var srcInt = int16(-255)
	// 请注意，之所以要执行uint16(srcInt)，是因为只有这样才能得到全二进制的表示。
	// 例如，fmt.Printf("%b", srcInt)将打印出"-11111111"，后者是负数符号再加上srcInt的绝对值的补码。
	// 而fmt.Printf("%b", uint16(srcInt))才会打印出srcInt原值的补码"1111111100000001"。
	fmt.Printf("The complement of srcInt is: %b (%b)\n", uint16(srcInt), srcInt)
	dstInt := int8(srcInt)
	fmt.Printf("The complement of srcInt is: %b (%b)\n", uint8(dstInt), dstInt)
	fmt.Printf("The value of dstInt is: %d\n", dstInt)
	fmt.Println()

	fmt.Printf("The Replacement Character is: %s\n", string(-1))
	fmt.Printf("The Unicode codepoint of Replacement Character: %U\n", '�')
	fmt.Println()

	srcStr := "你好"
	fmt.Printf("The string: %q\n", srcStr)
	fmt.Printf("The hex of %q: %x\n", srcStr, srcStr)
	fmt.Printf("The byte slice of %q: %x\n", srcStr, []byte(srcStr))
	fmt.Printf("The string: %q\n", string([]byte{'\xe4', '\xbd', '\xa0', '\xe5', '\xa5', '\xbd'}))
	fmt.Printf("The rune slice of %q: %U\n", srcStr, []rune(srcStr))
	fmt.Printf("The string %q\n", string([]rune{'\u4F60', '\u597D'}))

}
