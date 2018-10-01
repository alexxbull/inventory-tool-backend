package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

// retrieve the requested item from the database
func getItem(w http.ResponseWriter, req *http.Request) (Item, error) {
	// struct that will store all values from JSON data
	type TempItem struct {
		ID        string
		Name      string
		Price     string
		Color     string
		Condition string
		UserID    string
	}

	// store JSON data in struct object
	temp := TempItem{}
	err := json.NewDecoder(req.Body).Decode(&temp)

	if err != nil {
		http.Error(w, "incomplete input data", 400)
		return Item{}, err
	}

	// validate JSON values
	if temp.Name == "" || temp.Price == "" || temp.Color == "" || temp.Condition == "" || temp.UserID == "" {
		http.Error(w, "incomplete input data", 400)
		return Item{}, errors.New("Invalid item values")
	}

	// convert json values to primitive data types

	// convert item's price
	f64, err := strconv.ParseFloat(temp.Price, 64)
	if err != nil {
		http.Error(w, "invalid data", 400)
		return Item{}, err
	}
	itmPrice := f64

	// convert item's userID
	i64, err := strconv.Atoi(temp.UserID)
	if err != nil {
		http.Error(w, "invalid data", 400)
		return Item{}, err
	}
	itmUserID := i64

	// convert item's id
	var itmID int
	if temp.ID != "" {
		i64, err := strconv.Atoi(temp.ID)
		if err != nil {
			http.Error(w, "invalid data", 400)
			return Item{}, err
		}
		itmID = i64
	}

	// make proper Item object with correct data types
	item := Item{
		ID:        itmID,
		Name:      temp.Name,
		Price:     itmPrice,
		Color:     temp.Color,
		Condition: temp.Condition,
		UserID:    itmUserID,
	}

	return item, nil
}
