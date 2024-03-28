package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

var DB *sql.DB

func ConnectDatabase() error {
	db, err := sql.Open("sqlite3", "./personal.db")
	checkErr(err)

	query, err := DB.Query("CREATE TABLE IF NOT EXISTS customers ( customer_id INTEGER PRIMARY KEY, name TEXT,password TEXT );")
	checkErr(err)

	DB = db
	return nil
}

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

func main() {
	ConnectDatabase()
	getAllInfo()

}

func getAllInfo() Customer {
	rows, err := DB.Query("SELECT accounts.account_id, accounts.balance, customers.name FROM accounts JOIN customers ON accounts.customer_id = customers.customer_id;")
	checkErr(err)
	defer rows.Close()
	customer := Customer{}
	for rows.Next() {
		rows.Scan(&customer.ID, &customer.Name, &customer.Password)
	}
	return customer
}
