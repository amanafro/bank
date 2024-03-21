package db

import (
	"fmt"

	"github.com/google/uuid"
)

type account struct {
	ID      uuid    `json:"id"`
	Balance float32 `json:"balance"`
}

type customer struct {
	ID      uuid    `json:"c_id"`
	Name    string  `json:"name"`
	Account account `json:"account"`
}

func InitDB() {
	fmt.Println("3xp")
}
