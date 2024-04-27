package auth

import (
	"bank/db"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func LogIn() (bool, error) {
	db := db.InitDB()

	var userID int
	var password string
	var hash string

	fmt.Println("User ID")
	fmt.Scanln(&userID)
	fmt.Println("Passowrd")
	fmt.Scanln(&password)

	err := db.QueryRow("SELECT customer_id, password FROM customers WHERE customer_id = ? AND password = ?", userID).Scan(&hash)

	if err != nil {
		return false, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return false, nil
	}

	return true, nil
}

func Register() error {
	db := db.InitDB()

}
