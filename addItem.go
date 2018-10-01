package main

import (
	"net/http"
)

func addItem(w http.ResponseWriter, req *http.Request) {
	// check if request Method is valid
	if req.Method != "POST" {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	// get the item's properites
	item, err := getItem(w, req)
	if err != nil {
		return
	}

	// insert values into database
	sqlStatement := `
	INSERT INTO items (Name, Price, Color, Condition, UserID)
	VALUES ($1, $2, $3, $4, $5);`

	_, err = db.Exec(sqlStatement, item.Name, item.Price, item.Color, item.Condition, item.UserID)
	if err != nil {
		http.Error(w, "unable to insert data into database", 500)
	}
}
