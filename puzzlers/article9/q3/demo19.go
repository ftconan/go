// Author: magician
// File:   demo19.go
// Date:   2021/1/12
package main

import "fmt"

func main()  {
	var m map[string]int

	key := "two"
	elem, ok := m[key]
	fmt.Printf("The element pair with key %q in nil map: %d(%v)\n", key, elem, ok)
	fmt.Printf("The length of nil map: %d\n", len(m))
	fmt.Printf("Delete the key-element pair by key %q...\n", key)
	delete(m, key)

	m = map[string]int{}
	elem = 2
	fmt.Println("Add a key-element pair to a nil map...")
	m["two"] = elem
	fmt.Printf("m: %#v\n", m)

	//m1 := map[int]int{}
	//m1[2] = 2
	//fmt.Printf("m1: %#v\n", m1)
}
