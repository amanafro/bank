package transaction

import (
	"bank/db"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Account struct {
	ID         int
	CustomerID int
	Balance    float32
}

type Customer struct {
	ID       int
	Name     string
	Password string
}

func CheckError(err error) {
	if err != nil {
		log.Fatal("Error: ", err)
	}
}

const depositLimit float32 = 5.0

func Deposit() {
	db := db.InitDB()

	var depositeMoney float32

	fmt.Println("Great! How much you wanna deposit?")
	if _, err := fmt.Scanln(&depositeMoney); err != nil {
		log.Fatal("Error while depositing", err)
	}

	if depositeMoney > depositLimit {
		_, err := db.Exec("UPDATE account SET amount = amount + ? WHERE account.id=1 ", depositeMoney)
		CheckError(err)

		balance, err := db.Query("SELECT amount FROM account WHERE account.id=1")
		CheckError(err)

		for balance.Next() {
			var accountBalance Account
			err = balance.Scan(&accountBalance.Balance)
			CheckError(err)
			fmt.Println("Your deposit was succesful")
			fmt.Printf("Current account balance: %f\n", accountBalance.Balance)
		}
	} else {
		fmt.Println("You have to atleast deposit CHF 5")
	}
	db.Close()
}

const withdrawLimit float32 = 10.0

func Withdraw() {
	db := db.InitDB()

	var balance float32
	var withdrawMoney float32

	err := db.QueryRow("SELECT amount FROM account WHERE account.id=1").Scan(&balance)
	CheckError(err)

	fmt.Println("How much money would like to Withdraw? ")
	if _, err := fmt.Scanln(&withdrawMoney); err != nil {
		log.Fatal("Error while withdrawing the money")
	}

	if withdrawMoney > withdrawLimit {
		_, err := db.Exec("UPDATE account SET amount = amount - ? WHERE account.id=1", withdrawMoney)
		CheckError(err)

		balance, err := db.Query("SELECT amount FROM account WHERE account.id=1")
		CheckError(err)

		for balance.Next() {
			var accountBalance Account
			err = balance.Scan(&accountBalance.Balance)
			CheckError(err)

			fmt.Printf("Current account balance: %f\n", accountBalance.Balance)
		}
	}
	if withdrawMoney < withdrawLimit {
		fmt.Println("You can't withdraw less than CHF 10.-")
	}

	if withdrawMoney > balance {
		fmt.Println("Insufficient funds!")
	}
	db.Close()
}

func CheckBalance() {
	db := db.InitDB()

	update, err := db.Query("SELECT id, amount FROM account WHERE id=1 ")
	CheckError(err)

	for update.Next() {
		var accountBalance Account
		err = update.Scan(&accountBalance.ID, &accountBalance.Balance)
		CheckError(err)

		fmt.Printf("Account ID: %d\nAccount Balance: %f\n",
			accountBalance.ID, accountBalance.Balance)
	}
	db.Close()
}
