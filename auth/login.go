package auth

import (
	"bank/dbs"
	"bank/start"
	"database/sql"
	"errors"
	"fmt"
	"os"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/term"
)

type User struct {
	ID         int
	IsLoggedIn bool
}

var CurrentUser User

func ReadPassword(fd int) ([]byte, error) {
	return term.ReadPassword(fd)
}

func LogIn() (bool, error) {
	db, err := dbs.GetDB()
	dbs.CheckError(err)

	defer db.Close()

	var userID int
	var hashedPassword string

	fmt.Println("User ID")
	fmt.Scanln(&userID)
	var password string

	var passwordBytes []byte

	fmt.Println("Password")
	passwordBytes, err = ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		return false, fmt.Errorf("Error reading password: %v", err)
	}
	password = string(passwordBytes)

	if len(password) == 0 {
		return false, fmt.Errorf("Password cannot be empty")
	}

	err = db.QueryRow("SELECT id, password_hash FROM account WHERE id = ?", userID).Scan(&userID, &hashedPassword)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, fmt.Errorf("User not found")
		}
		return false, fmt.Errorf("Error querying database: %v", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return false, fmt.Errorf("Invalid password")
	} else {
		fmt.Println("You've succesfully logged in")
		start.Intro()
	}
	return true, nil
}
