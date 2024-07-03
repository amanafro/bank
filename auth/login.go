package auth

import (
	"bank/dbs"
	"bank/start"
	"database/sql"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/term"
	"os"
)

func ReadPassword(fd int) ([]byte, error) {
	return term.ReadPassword(fd)
}

func LogIn() (bool, error) {
outer:
	db, err := dbs.GetDB()
	if err != nil {
		return false, &loginError{msg: fmt.Sprintf("Error connecting to database: %v", err)}
	}
	defer db.Close()

	fmt.Print("User ID: ")
	var userID int
	if _, err := fmt.Scanln(&userID); err != nil {
		return false, &loginError{msg: "Error reading user ID"}
	}

	passwordBytes, err := ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		return false, &loginError{msg: fmt.Sprintf("Error reading password: %v", err)}
	}
	password := string(passwordBytes)

	if len(password) == 0 {
		fmt.Println("Password cannot be empty")
		return false, nil
	}

	row := db.QueryRow("SELECT id, password FROM account WHERE id = ?", userID)
	if err := row.Scan(&userID, &hashedPassword); err != nil {
		switch err {
		case sql.ErrNoRows:
			fmt.Println("User not found")
			break outer
		default:
			return false, &loginError{msg: fmt.Sprintf("Error querying database: %v", err)}
		}
	}

	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		fmt.Println("Invalid password")
		return false, nil
	} else {
		fmt.Println("You've successfully logged in")
		start.Intro()
	}

	return true, nil
}

type loginError struct {
	msg string
}
