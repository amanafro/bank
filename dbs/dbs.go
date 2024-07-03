package dbs

import (
	"database/sql"
	"log"
)

func GetDB() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/bank")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec(`
  CREATE TABLE account (
    id INT PRIMARY KEY,
    name VARCHAR(255),
    email VARCHAR(255) NOT NULL,
    address VARCHAR(255),
)`)
	if err != nil {
		panic(err)
	}
}

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
