package transaction

import (
	"bank/db"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Account struct {
	ID         int
	Number     string
	Balance    int
	CustomerID int
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

const depositLimit int = 5

func Deposit() {
	db := db.InitDB()

	var depositeMoney int

	fmt.Println("Great! How much you wanna deposit?")
	if _, err := fmt.Scanln(&depositeMoney); err != nil {
		log.Fatal("Error while depositing", err)
	}

	if depositeMoney > depositLimit {
		_, err := db.Exec("UPDATE accounts SET balance = balance + ? WHERE account_id=1 ", depositeMoney)
		CheckError(err)

		balance, err := db.Query("SELECT balance FROM accounts WHERE account_id=1")
		CheckError(err)

		for balance.Next() {
			var accountBalance Account
			err = balance.Scan(&accountBalance.Balance)
			CheckError(err)
			fmt.Println("Your deposit was succesful")
			fmt.Printf("Current account balance: %d\n", accountBalance.Balance)
		}
	} else {
		fmt.Println("You have to atleast deposit CHF 5")
	}
	db.Close()
}

const withdrawLimit int = 10

func Withdraw() {
	db := db.InitDB()

	var balance int
	var withdrawMoney int

	err := db.QueryRow("SELECT balance FROM accounts WHERE account_id=1").Scan(&balance)
	CheckError(err)

	fmt.Println("How much money would like to Withdraw? ")
	if _, err := fmt.Scanln(&withdrawMoney); err != nil {
		log.Fatal("Error while withdrawing the money")
	}

	if withdrawMoney < withdrawLimit {
		update, err := db.Exec("UPDATE accounts SET balance = balance - ? WHERE account_id=1", withdrawMoney)
		CheckError(err)

		balance, err := db.Query("SELECT balance FROM accounts WHERE account_id=1")
		CheckError(err)

		for balance.Next() {
			var accountBalance Account
			err = balance.Scan(&accountBalance.Balance)
			CheckError(err)

			fmt.Printf("Current account balance: %d\n", accountBalance.Balance)
		}
		fmt.Println(&update)
	}
	if withdrawMoney > balance {
		fmt.Println("Insufficient funds!")
	} else {
		fmt.Println("You can't withdraw less than CHF 10")
	}
	db.Close()
}

func CheckBalance() {
	db := db.InitDB()

	update, err := db.Query("SELECT account_id, account_number, balance FROM accounts WHERE customer_id=1 ")
	CheckError(err)

	for update.Next() {
		var accountBalance Account
		err = update.Scan(&accountBalance.ID, &accountBalance.Number, &accountBalance.Balance)
		CheckError(err)

		fmt.Printf(" Account ID: %d,\n Account Num: %s,\n Account Balance: %d\n",
			accountBalance.ID, accountBalance.Number, accountBalance.Balance)
	}
	db.Close()
}
