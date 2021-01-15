// Author: magician
// File:   demo31.go
// Date:   2021/1/15
package main

import "fmt"

type Pet interface {
	SetName(name string)
	Name() string
	Category() string
}

type Dog struct {
	name string
}

func (dog *Dog) SetName(name string) {
	dog.name = name
}

func (dog Dog) Name() string {
	return dog.name
}

func (dog Dog) Category() string {
	return "dog"
}

func main() {
	dog := Dog{"dog"}
	_, ok := interface{}(dog).(Pet)
	fmt.Printf("Dog implements interface Pet: %v\n", ok)
	_, ok = interface{}(&dog).(Pet)
	fmt.Printf("Dog implements interface Pet: %v\n", ok)
	fmt.Println()

	var pet Pet = &dog
	fmt.Printf("This pet is a %s, the name is %q.\n", pet.Category(), pet.Name())
}
