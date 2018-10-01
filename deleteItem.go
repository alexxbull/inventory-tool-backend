package main

import (
	"net/http"
	"strconv"
)

func deleteItem(w http.ResponseWriter, req *http.Request) {
	// check if request Method is valid
	if req.Method != "DELETE" {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	// get id of item to be deleted
	id := req.Header.Get("ID")

	// validate id
	if id == "" {
		http.Error(w, http.StatusText(400), 400)
		return
	}

	// convert id value
	itemID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	// delete item from database
	sqlStatement := `DELETE FROM items WHERE id = $1`

	_, err = db.Exec(sqlStatement, itemID)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
	}

}
