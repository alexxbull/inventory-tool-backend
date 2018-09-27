package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

func getItem(w http.ResponseWriter, req *http.Request) (Item, error) {
	// struct that will store all values from JSON data
	type tempItem struct {
		ID        string
		Name      string
		Price     string
		Color     string
		Condition string
		UserID    string
	}

	// read JSON data into struct
	temp := tempItem{}
	err := json.NewDecoder(req.Body).Decode(&temp)

	if err != nil {
		http.Error(w, http.StatusText(400), 400)
		return Item{}, err
	}

	// validate JSON values
	if temp.Name == "" || temp.Price == "" || temp.Color == "" || temp.Condition == "" || temp.UserID == "" {
		http.Error(w, http.StatusText(400), 400)
		return Item{}, errors.New("Invalid item values")
	}

	// convert json values
	f64, err := strconv.ParseFloat(temp.Price, 64)
	if err != nil {
		http.Error(w, http.StatusText(406), 406)
		return Item{}, err
	}
	itmPrice := f64

	i64, err := strconv.Atoi(temp.UserID)
	if err != nil {
		http.Error(w, http.StatusText(406), 406)
		return Item{}, err
	}
	itmUserID := i64

	var itmID int
	if temp.ID != "" {
		i64, err := strconv.Atoi(temp.ID)
		if err != nil {
			http.Error(w, http.StatusText(406), 406)
			return Item{}, err
		}
		itmID = i64
	}

	// make proper item
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
