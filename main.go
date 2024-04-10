package main

import (
	"database/sql"
	"fmt"
	"log"

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

func InitDB() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/gobank")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	return db
}

func main() {
	intro()
}

func intro() {
	var choice string

	fmt.Println("Hello! How can help you? \n 1. Deposit \n 2. Withdraw \n 3. Balance")
	fmt.Scanln(&choice)

	switch choice {
	case "1":

	}

}

func Deposit() {
	db := InitDB()

	var depositeMoney int

	fmt.Println("Great! How much you wanna deposit?")
	fmt.Scanln(&depositeMoney)

	if depositeMoney < 5 {
		fmt.Println("You have to atleast deposit CHF 5")
	}

	res, err := db.Query("UPDATE accounts SET balance+ WHERE account_id='1' ")
  if err != nil {
    log.Fatal(err)
  }

  res.Next() {
    for 
  }


}
