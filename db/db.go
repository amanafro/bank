package db

import (
	"database/sql"
)

func InitDB() *sql.DB {
	db, err := sql.Open("mysql", "aman:root@tcp(127.0.0.1:3306)/bank")
	if err != nil {
		panic(err)
	}
	return db
}
