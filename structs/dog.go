package main

import "fmt"

type Animal interface {
	Name() string
}

type Dog struct{}

func (d *Dog) Name() string {
	return "Dog"
}

func (d *Dog) Bark() string {
	return "Woof!"
}

func main() {
	var animal Animal
	animal = &Dog{}
	fmt.Println("That animal is a", animal.Name())
}
