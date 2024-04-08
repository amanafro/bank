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

	result, err := db.Query("SELECT a.account_id, a.account_number, a.balance, c.name AS customer_name FROM   accounts a  JOIN customers c ON a.customer_id = c.customer_id;")
	if err != nil {
		panic(err)
	}

	for result.Next() {
		var account Account
		var cusomer Customer

		err = result.Scan(&cusomer.ID, &cusomer.Name, &account.CustomerID, &account.Balance)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Customer: %s, CustomerID: %d, Balance: %f", cusomer.Name, account.CustomerID, account.Balance)
	}
}

func intro() {
	fmt.Println("Hello! How can help you")
	fmt.Printf("1. H")

}
