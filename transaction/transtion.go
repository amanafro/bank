package transaction

import (
	"fmt"
	"log"

	"bank/db"

	_ "github.com/go-sql-driver/mysql"
)

type Customer struct {
	ID       int
	Name     string
	Email    string
	Password string
	Balance  float32
}

func CheckError(err error) {
	if err != nil {
		log.Fatal("Error: ", err)
	}
}

const depositLimit float32 = 5.0

func Deposit() {
	db, err := db.GetDB()
	if err != nil {
		fmt.Println("Error getting DB connection:", err)
	}
	defer db.Close()

	var depositeMoney float32

	fmt.Println("Great! How much you wanna deposit?")
	if _, err := fmt.Scanln(&depositeMoney); err != nil {
		log.Fatal("Error while depositing", err)
	}

	var balance float32
	err = db.QueryRow("SELECT balance FROM account WHERE id=1").Scan(&balance)
	if err != nil {
		log.Fatal(err)
	}

	if depositeMoney < depositLimit {
		fmt.Println("You need to atleast deposit CHF 5.-")
		return
	}

	_, err = db.Exec("UPDATE account SET balance = balance + ? WHERE account.id = 1", depositeMoney)
	CheckError(err)

	updatedBalance := balance + depositeMoney

	fmt.Printf("Transaction succesful. \n Current account balance: %f\n", updatedBalance)
}

const withdrawLimit float32 = 10.0

func Withdraw() {
	db, err := db.GetDB()
	if err != nil {
		fmt.Println("Error getting DB connection:", err)
		return
	}
	defer db.Close()

	var balance float32
	err = db.QueryRow("SELECT balance FROM account WHERE id=1").Scan(&balance)
	if err != nil {
		log.Fatal(err)
	}

	var withdrawMoney float32

	fmt.Println("How much money would like to Withdraw? ")
	if _, err := fmt.Scanln(&withdrawMoney); err != nil {
		log.Fatal("Error while withdrawing the money")
	}
	if withdrawMoney > balance {
		fmt.Println("Insufficient funds!")
		return
	}
	if withdrawMoney < withdrawLimit {
		fmt.Println("You can't withdraw less than CHF 10.-")
		return
	}
	_, err = db.Exec("UPDATE account SET balance = balance - ? WHERE account.id = 1", withdrawMoney)
	CheckError(err)

	updatedBalance := balance + withdrawMoney

	fmt.Printf("Transaction succesful. \n Current account balance: %f\n", updatedBalance)
}

func CheckBalance() {
	db, err := db.GetDB()
	if err != nil {
		fmt.Println("Error getting DB connection:", err)
		return
	}
	defer db.Close()

	update, err := db.Query("SELECT id, balance FROM account WHERE id=1")
	CheckError(err)

	for update.Next() {
		var accountBalance Customer

		err = update.Scan(&accountBalance.ID, &accountBalance.Balance)
		CheckError(err)

		fmt.Printf("Account ID: %d\nAccount Balance: %f\n",
			accountBalance.ID, accountBalance.Balance)
	}
}
