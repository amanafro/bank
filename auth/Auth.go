package auth

import (
	"fmt"
)

func Auth() {
	var choice string

	fmt.Println("Login or Register (log/reg)")
	fmt.Scanln(&choice)

	switch choice {
	case "log":

		fmt.Println("Welcome back")
		LogIn()

	case "reg":

		fmt.Println("New here? Create an account")
		Register()
	}
}
