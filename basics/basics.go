package main

import "fmt"

func main() {
	fmt.Println("One + One = ", 2) // One + One = 2
	fmt.Println(true || false)     // true

	var x = "a"
	y := 1
	fmt.Println(x) // a
	fmt.Println(y) // 1
	var a, b, c int = 1, 2, 3
	fmt.Println(a + b + c) // 6
}
