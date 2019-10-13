package main

import "fmt"

func main() {
	for i := 1; i < 10; i++ {
		fmt.Println(i)
	}
	for {
		fmt.Println("infinite loop")
		break
	}

	for index, value := range "loop thru string" {
		fmt.Printf("char at %d = %q\n", index, value)
	}

	nums := []int{1, 2, 3}
	for _, val := range nums {
		fmt.Println(val)
	}
}
