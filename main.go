package main

import (
	"bank/auth"
	"bank/transaction"
	"fmt"
)

func main() {
	auth.Register()

}

func intro() {
	var choice string

	fmt.Println("Hello! How can help you? \n 1. Deposit \n 2. Withdraw \n 3. Balance")
	fmt.Scanln(&choice)

	switch choice {
	case "1":
		transaction.Deposit()
	case "2":
		transaction.Withdraw()
	case "3":
		transaction.CheckBalance()
	default:
		fmt.Println("Invalid choice")
	}
}
