package auth

import (
	"bank/db"
	"fmt"
	"log"

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

func Register() {
	db := db.InitDB()

	var customer_name string
	var customer_password string
	var customer_password_conform string

	fmt.Printf("Register \n")
	fmt.Println("Full name")
	fmt.Scanln(&customer_name)
	fmt.Println("Passowrd")
	fmt.Scanln(&customer_password)
	fmt.Println("Conform passowrd")
	fmt.Scanln(&customer_password_conform)

	if customer_password == customer_password_conform {

		hash, err := bcrypt.GenerateFromPassword([]byte(customer_password), bcrypt.DefaultCost)
		if err != nil {
			log.Fatal(err)
		}

		_, err = db.Exec("INSERT INTO customers (name, password) VALUES (?,?)", customer_name, hash)

		log.Fatal(err)

	} else {
		fmt.Println("The passwords dont much")
	}
}
