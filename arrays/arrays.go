package main

import "fmt"

func main() {
	var a [10]int
	fmt.Println("empty: ", a)
	a[3] = 25
	fmt.Println("get:", a[3])

	b := []int{1, 2, 3, 4, 5}
	fmt.Println(b)
}
