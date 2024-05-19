package transaction

import (
	"bank/dbs"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Customer struct {
	ID       int
	Name     string
	Email    string
	Password string
	Balance  float32
}

// depositing

const depositLimit float32 = 5.0

func Deposit() (float32, error) {
	db, err := dbs.GetDB()
	if err != nil {
		fmt.Println("Error getting DB connection:", err)
		return 0, err
	}
	defer db.Close()

	var depositeMoney float32

	fmt.Println("Great! How much you wanna deposit?")
	if _, err := fmt.Scanln(&depositeMoney); err != nil {
		log.Fatal("Error while depositing", err)
	}

	var balance float32
	err = db.QueryRow("SELECT balance FROM accounts WHERE id=1").Scan(&balance)
	if err != nil {
		return 0, fmt.Errorf("error fetching balance: %v", err)
	}

	if depositeMoney < depositLimit {
		fmt.Println("You need to atleast deposit CHF 5.-")
		return 0, err
	}

	_, err = db.Exec("UPDATE account SET balance = balance - ? WHERE account.id = 1", depositeMoney)
	dbs.CheckError(err)

	fmt.Printf("Transaction succesful. \n Current account balance: %f\n", balance)

	return balance, err
}

// withdrawing

const withdrawLimit float32 = 10.0

func Withdraw() {
	db, err := dbs.GetDB()
	if err != nil {
		fmt.Println("Error getting DB connection:", err)
		return
	}
	defer db.Close()

	var balance float32
	var withdrawMoney float32

	err = db.QueryRow("SELECT balance FROM account WHERE account.id=1").Scan(balance)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("How much money would like to Withdraw? ")
	if _, err := fmt.Scanln(&withdrawMoney); err != nil {
		log.Fatal("Error while withdrawing the money")
		return
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
	dbs.CheckError(err)

	fmt.Printf("Transaction succesful. \n Current account balance: %f\n", balance)

}

// balance checking

func CheckBalance() {
	db, err := dbs.GetDB()
	if err != nil {
		fmt.Println("Error getting DB connection:", err)
		return
	}
	defer db.Close()

	update, err := db.Query("SELECT id, balance FROM account WHERE id=1")
	dbs.CheckError(err)

	for update.Next() {
		var accountBalance Customer

		err = update.Scan(&accountBalance.ID, &accountBalance.Balance)
		dbs.CheckError(err)

		fmt.Printf("Account ID: %d\nAccount Balance: %f\n",
			accountBalance.ID, accountBalance.Balance)
	}
}
