package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var db *sql.DB

func main() {
	InitDB()

	getInfo1()
}

type Account struct {
	ID      int     `json:"id"`
	Balance float32 `json:"balance"`
}

type Customer struct {
	ID      int     `json:"c_id"`
	Name    string  `json:"name"`
	Account Account `json:"account"`
}

func errCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func InitDB() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/personal")
	defer db.Close()

	errCheck(err)
}

func getInfo1() {
	rows, err := db.Query("SELECT * FROM account, customer where ")
	errCheck(err)
	defer rows.Close()

	var info Account
	for rows.Next() {
		var account Account
		err = rows.Scan(&account.ID, &account.Balance)
		errCheck(err)
		log.Printf(info.ID)
	}
}
