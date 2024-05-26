package dbs

import (
	"database/sql"
	"log"
)

func GetDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/bank")
	if err != nil {
		return nil, err
	}
	return db, nil
}

func CheckError(err error) {
	if err != nil {
		log.Fatal("Error: ", err)
	}
}
