// Author: magician
// File:   demo36.go
// Date:   2021/1/15
package main

import "fmt"

type Dog struct {
	name string
}

func New(name string) Dog {
	return Dog{name}
}

func (dog *Dog) SetName(name string) {
	dog.name = name
}

func (dog Dog) Name() string {
	return dog.name
}

func main()  {
	//New("little pig").SetName("monster") // 不能调用不可寻址的值的指针方法。

	map[string]int{"the": 0, "word": 0, "counter": 0}["word"]++
	map1 := map[string]int{"the": 0, "word": 0, "counter": 0}
	map1["word"]++
	fmt.Printf("%v\n", map1)
}
