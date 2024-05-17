package auth

import (
	"fmt"
)

func Auth() {
	var choice string

	fmt.Println("Welcome, Login (log)")
	fmt.Scanln(&choice)

	fmt.Println("New here? Create an account (reg)")
	fmt.Scanln(&choice)

	switch choice {
	case "log":
		LogIn()
	case "reg":
		Register()
	}
}
