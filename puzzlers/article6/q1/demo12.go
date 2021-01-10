// Author: magician
// File:   demo12.go
// Date:   2021/1/10
package main

import "fmt"

var container = []string{"zero", "one", "two"}

func main() {
	container := map[int]string{0: "zero", 1: "one", 2: "two"}
	_, ok1 := interface{}(container).([]string)
	_, ok2 := interface{}(container).(map[int]string)

	if !(ok1 || ok2) {
		fmt.Printf("Error: unsupported container type: %T\n", container)
		return
	}
	fmt.Printf("The element is: %q. (the container type is %T)\n", container[1], container)

	elem, err := getElement(container)
	if err != nil {
		fmt.Printf("Error is %s\n", err)
		return
	}
	fmt.Printf("The element is: %q. (the container type is %T)\n", elem, container)
}

func getElement(containerI interface{}) (elem string, err error) {
	switch t := containerI.(type) {
	case []string:
		elem = t[1]
	case map[int]string:
		elem = t[1]
	default:
		err = fmt.Errorf("unsupported container type: %T\n", containerI)
		return
	}
	return
}
