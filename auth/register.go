package auth

import (
	"bank/dbs"
	"bank/start"
	"fmt"
	"log"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func Register() {
	db, err := dbs.GetDB()
	if err != nil {
		fmt.Println("Error getting DB connection:", err)
	}
	defer db.Close()

	var customer_name string

	var customer_email string

	var customer_password []byte
	var customer_password_conform []byte

	fmt.Printf("Register \n")
	fmt.Printf("Full name \n")
	fmt.Scanln(&customer_name)
	fmt.Printf("Email \n")
	fmt.Scanln(&customer_email)
	fmt.Println("Password")

	customer_password, err = ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		log.Fatalf("Error reading password: %v", err)
	}
	password := string(customer_password)

	if len(password) == 0 {
		log.Fatal("Password cannot be empty")
	}

	fmt.Println("Confirm password")

	customer_password_conform, err = ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		log.Fatalf("Error reading password: %v", err)
	}
	confirm_password := string(customer_password_conform)

	if len(confirm_password) == 0 {
		log.Fatal("Password cannot be empty")
	}

	if confirm_password == password {
		hash, err := bcrypt.GenerateFromPassword([]byte(customer_password), bcrypt.DefaultCost)
		if err != nil {
			log.Fatal(err)
		}
		_, err = db.Exec("INSERT INTO account (name, email, password, balance) VALUES (?,?,?,?)", customer_name, customer_email, hash, 0)
		log.Fatal(err)
		fmt.Println("You have been registered successfully")
		start.Intro()
	} else {
		fmt.Println("The passwords dont much")
	}
}
