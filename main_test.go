package main

import (
	"database/sql"
	"fmt"
	"os"
	"testing"
	"github.com/stretchr/testify/assert"
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
	assert.NotNil(t, food)
	assert.Equal(t, 1, food.id, "Food ID 1 should be equal")
	food = readFood(3)
	assert.NotNil(t, food)
	assert.Equal(t, 3, food.id, "Food ID 3 should be equal")
	food = readFood(-1)
	assert.Nil(t, food)
}

func TestInsertFood(t *testing.T) {
	food := Food{4, "Spagetthi", 12000, "La Fonte"}
	lastInsertId := insertFood(food)
	assert.NotEqual(t, 0, lastInsertId, "Food should have been inserted")

	insertedFood := readFood(lastInsertId)
	assert.NotNil(t, insertedFood)
}

func TestDeleteFood(t *testing.T) {
	foodInserted := Food{0, "Kopi Aku Kamu", 18000, "Aku Kamu"}
	lastInsertId := insertFood(foodInserted)
	assert.NotEqual(t, 0, lastInsertId, "Food should have been inserted")

	affectedRows := deleteFood(lastInsertId)
	assert.Equal(t, 1, int(affectedRows), "There should be only 1 food affected")
}

func TestUpdateFoodPrice(t *testing.T) {
	affectedRows := updateFoodPrice(1, 28000)
	assert.Equal(t, 1, affectedRows, "There should be only 1 food affected")
	food := readFood(1)
	assert.Equal(t, 1, food.id, "Effects of update should have applied only to food ID 1")
	assert.Equal(t, 28000, food.price, "Food price should have been updated")
}
