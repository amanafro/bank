package auth

import (
	"fmt"
)

func Auth() {
	var choice string

	fmt.Println("Login or Register (1/2)")
	fmt.Scanln(&choice)

	switch choice {
	case "1":

		fmt.Println("Welcome back")
		LogIn()

	case "2":

		fmt.Println("New here? Create an account")
		Register()
	}
}
