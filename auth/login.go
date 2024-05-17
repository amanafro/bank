package auth

import (
	"bank/db"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func LogIn() (bool, error) {

	db, err := db.GetDB()
	if err != nil {
		fmt.Println("Error getting DB connection:", err)
	}
	defer db.Close()

	var userID int
	var password string
	var hash string

	fmt.Println("User ID")
	fmt.Scanln(&userID)
	fmt.Println("Passowrd")
	fmt.Scanln(&password)

	err = db.QueryRow("SELECT id, password FROM account WHERE id = ?", userID).Scan(&hash)
	if err != nil {
		log.Fatal(err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return false, nil
	}

	return true, nil
}
