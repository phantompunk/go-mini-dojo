package main

import "fmt"

const greeting = "Hello "

func Hello(name string) string {
	if name == "" {
		name = "you"
	}
	return greeting + name
}

func main() {
	fmt.Println(Hello("Rigo"))
}
