package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Customer struct {
	ID       int
	Name     string
	Password string
}

type Account struct {
	ID         int
	Balance    float64
	CustomerID int
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/gobank")
	checkErr(err)

	defer db.Close()

	result, err := db.Query("SELECT a.account_id, a.account_number, a.balance, c.name AS customer_name FROM   accounts a  JOIN customers c ON a.customer_id = c.customer_id;")
	checkErr(err)

	for result.Next() {
		var cusomer Customer
		var account Account

		err = result.Scan(&cusomer.Name, &account.CustomerID, &account.Balance)
		checkErr(err)
		fmt.Printf("Customer: %s, CustomerID: %d, Balance: %f", cusomer.Name, account.CustomerID, account.Balance)
	}
}

func intro() {
	fmt.Println("Hello! How can help you")
	fmt.Printf("1. H")

}
