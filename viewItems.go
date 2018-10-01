package main

import (
	"encoding/json"
	"net/http"

	_ "github.com/lib/pq"
)

func viewItems(w http.ResponseWriter, req *http.Request) {
	// check if request Method is valid
	if req.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	// get the user's ID from the header
	userID := req.Header.Get("UserID")
	if userID == "" {
		http.Error(w, "user id not found", 400)
		return
	}

	// query database for all items belonging to user
	sqlGetItems := `
		SELECT * 
		FROM items
		WHERE UserID = $1`

	rows, err := db.Query(sqlGetItems, userID)
	if err != nil {
		http.Error(w, "unable to select records from database", 500)
		return
	}
	defer rows.Close()

	// store all items from query in a slice
	items := make([]Item, 0)
	for rows.Next() {
		item := Item{}
		err = rows.Scan(&item.ID, &item.Name, &item.Price, &item.Color, &item.Condition, &item.UserID)
		if err != nil {
			http.Error(w, "unable to retrieve record from database", 500)
			return
		}
		items = append(items, item)
	}

	// check if rows encountered any errors
	if err = rows.Err(); err != nil {
		http.Error(w, "unable to retrieve records from database", 500)
		return
	}

	// send the queried items as JSON
	jsonItems, err := json.Marshal(items)
	if err != nil {
		http.Error(w, "unable to respond with retrieved records", 500)
		return
	}
	w.Write(jsonItems)
}
