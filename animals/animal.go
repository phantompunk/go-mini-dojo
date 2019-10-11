package main

import "fmt"

type Runner interface {
	Run() string
}

type Animal struct {
	Name string
	Age  int
	Legs int
}

type Dog struct {
	Animal
}

func (d *Dog) Run() string {
	return "I'm a runner!"
}

func main() {
	dog := Dog{
		Animal{
			Name: "Champion",
			Age:  2,
			Legs: 3,
		},
	}
	fmt.Println(dog.Run())
}
