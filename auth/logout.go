package auth

import (
	"bank/dbs"
	"fmt"
)

func Logout() {
	db, err := dbs.GetDB()
	dbs.CheckError(err)

	_, err = db.Exec("UPDATE account SET is_logged_in = false WHERE id = ?", CurrentUser.ID)
	if err != nil {
		fmt.Println("Error logging out:", err)
		return
	}

	CurrentUser = User{} // Reset the CurrentUser
	fmt.Println("Logged out successfully!")
}
