package main

import (
	"net/http"
)

func editItem(w http.ResponseWriter, req *http.Request) {
	// check if request Method is valid
	if req.Method != "PUT" {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	// get the item's properites
	item, err := getItem(w, req)
	if err != nil {
		return
	}

	// UPDATE Item
	sqlStatement := `
	UPDATE items
	SET Name = $1, Price = $2, Color = $3, Condition = $4
	WHERE ID = $5;`

	_, err = db.Exec(sqlStatement, item.Name, item.Price, item.Color, item.Condition, item.ID)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
	}
}
