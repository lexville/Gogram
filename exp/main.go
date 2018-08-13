package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host   = "localhost"
	port   = 5432
	user   = "postgres"
	dbname = "gogram_dev"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	for i := 1; i <= 6; i++ {
		userID := 1
		if i > 3 {
			userID = 2
		}
		amount := i * 100
		description := fmt.Sprintf("USB-C Adapter x%d", i)
		_, err = db.Exec(`
		INSERT INTO orders(user_id, amount, description)
		VALUES($1, $2, $3)`, userID, amount, description)
		if err != nil {
			panic(err)
		}
	}

	rows, err := db.Query(`
		SELECT *
		FROM users
		INNER JOIN orders on users.id=orders.user_id
	`)

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var userID, orderID, amount int
		var email, name, desc string
		if err := rows.Scan(&userID, &name, &email, &orderID, &userID, &amount, &desc); err != nil {
			panic(err)
		}
		fmt.Println("userID:", userID, "name:", name, "email:", email, "orderID:", orderID, "amount:", amount, "desc:", desc)
	}

	if rows.Err() != nil {
		panic(rows.Err())
	}
}
