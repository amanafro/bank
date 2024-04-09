package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Account struct {
	ID         int
	Balance    int
	CustomerID int
}

type Customer struct {
	ID       int
	Name     string
	Password string
}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/gobank")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	result, err := db.Query("SELECT a.account_id, c.name, a.balance FROM  accounts a  JOIN customers c ON a.customer_id = c.customer_id;")
	if err != nil {
		panic(err)
	}

	for result.Next() {
		var account Account
		var customer Customer

		err = result.Scan(&account.ID, &customer.Name, &account.Balance)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Customer: %s, AccountID: %d, Balance: %d \n", customer.Name, account.ID, account.Balance)
	}
}

func intro() {
	fmt.Println("Hello! How can help you")
	fmt.Printf("1. H")

}
