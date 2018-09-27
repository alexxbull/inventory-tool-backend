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

	// query database for all items
	rows, err := db.Query("SELECT * from items;")
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	defer rows.Close()

	// store all items from query in a slice
	items := make([]Item, 0)
	for rows.Next() {
		item := Item{}
		err = rows.Scan(&item.ID, &item.Name, &item.Price, &item.Color, &item.Condition, &item.UserID)
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}
		items = append(items, item)
	}

	if err = rows.Err(); err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	// send the queried items as JSON
	jsonItems, err := json.Marshal(items)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	w.Write(jsonItems)

}
