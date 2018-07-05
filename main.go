package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Food struct {
	id    int
	name  string
	price int
	owner string
}

const (
	DB_USERNAME = "postgres"
	DB_PASSWORD = "postgres"
	DB_NAME     = "food_test"
)

var (
	db  *sql.DB
	err error
)

func main() {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USERNAME, DB_PASSWORD, DB_NAME)
	db, err = sql.Open("postgres", dbinfo)
	checkError(err)
	defer db.Close()
}

func readFood(id int) *Food {
	row := db.QueryRow("SELECT * FROM food WHERE ID=$1", id)

	var foodId int
	var name string
	var price int
	var owner string

	err = row.Scan(&foodId, &name, &price, &owner)
	if err == sql.ErrNoRows {
		return nil
	} else {
		checkError(err)
		return &Food{foodId, name, price, owner}
	}
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
