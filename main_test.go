package main

import (
	"database/sql"
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USERNAME, DB_PASSWORD, DB_NAME)
	db, err = sql.Open("postgres", dbinfo)
	checkError(err)
	defer db.Close()

	testResult := m.Run()

	os.Exit(testResult)
}

func TestReadFood(t *testing.T) {
	food := readFood(1)
	if food == nil {
		t.Errorf("Food 1 is nil")
	}
	if food.id != 1 {
		t.Errorf("Food ID is incorrect")
	}
	food = readFood(3)
	if food == nil {
		t.Errorf("Food 3 is nil")
	}
	if food.id != 3 {
		t.Errorf("Food 3 ID is incorrect")
	}
	food = readFood(-1)
	if food != nil {
		t.Errorf("Food 5 is not nil")
	}
}

func TestInsertFood(t *testing.T) {
	food := Food{4, "Spagetthi", 12000, "La Fonte"}
	result := insertFood(food)

	if result == 0 {
		t.Errorf("Food is nil")
	}

	expectedFood := readFood(result)

	if expectedFood == nil {
		t.Errorf("Food not inserted")
	}
}

func TestDeleteFood(t *testing.T) {
	foodInserted := Food{0, "Kopi Aku Kamu", 18000, "Aku Kamu"}
	lastInsertId := insertFood(foodInserted)

	if lastInsertId == 0 {
		t.Errorf("Food should have been inserted")
	}

	affectedRows := deleteFood(lastInsertId)
	if affectedRows != 1 {
		t.Errorf("Food should have been only one")
	}
}
