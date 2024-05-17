package auth

import (
	"bank/db"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func Register() {
	db, err := db.GetDB()
	if err != nil {
		fmt.Println("Error getting DB connection:", err)
	}
	defer db.Close()

	var customer_name string
	var customer_email string
	var customer_password string
	var customer_password_conform string

	fmt.Printf("Register \n")
	fmt.Printf("Full name \n")
	fmt.Scanln(&customer_name)
	fmt.Printf("Email \n")
	fmt.Scanln(&customer_email)
	fmt.Printf("Passowrd \n")
	fmt.Scanln(&customer_password)
	fmt.Printf("Conform passowrd \n")
	fmt.Scanln(&customer_password_conform)

	if customer_password == customer_password_conform {

		hash, err := bcrypt.GenerateFromPassword([]byte(customer_password), bcrypt.DefaultCost)
		if err != nil {
			log.Fatal(err)
		}
		_, err = db.Exec("INSERT INTO account (name, email, password) VALUES (?,?)", customer_name, hash)
		log.Fatal(err)

	} else {
		fmt.Println("The passwords dont much")
	}
}
