package main

import (
	"errors"
	"fmt"
	"log"
)

func WithdrawAmount(balance, amount int) (int, error) {
	if amount > balance {
		return -1, errors.New("Insufficient funds")
	}
	return balance - amount, nil
}
func main() {
	func() {
		fmt.Println("Hello from anon fun")
	}()

	bal, err := WithdrawAmount(100, 30)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Balance after withdrawl: ", bal)
}
