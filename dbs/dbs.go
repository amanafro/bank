package dbs

import (
	"database/sql"
	"log"
)

func initializeDB() error {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/bank")
	if err != nil {
		return err
	}

	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS account (
        id INT AUTO_INCREMENT PRIMARY KEY,
        name VARCHAR(255) NOT NULL,
        email VARCHAR(255) UNIQUE NOT NULL,
        password_hash VARCHAR(255) NOT NULL,
        balance DECIMAL(10, 2) NOT NULL DEFAULT 0.00,
    )`)
	if err != nil {
		return err
	}

	db.Close() // Close after initialization
	return nil
}

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
