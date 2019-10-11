package main

import "fmt"

func main() {
	s := make([]string, 2)
	fmt.Println("empty:", s)
	s[0], s[1] = "zero", "one"
	fmt.Println("set:", s)

	s = append(s, "two")
	s = append(s, "three")
	ex := s[:2]
	fmt.Println(ex)
	in := s[1:]
	fmt.Println(in)
}
