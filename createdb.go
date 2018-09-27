package main

import (
	"log"
	"net/http"
)

// create a database table if one doesn't exist
func createDBT() error {
	sqlStatement := `
	CREATE TABLE IF NOT EXISTS inventory.items(
		ID Serial PRIMARY KEY NOT NULL,
		Name TEXT  NOT NULL,
		Price NUMERIC(11, 2) NOT NULL,
		Color TEXT NOT NULL NOT NULL,
		Condition VARCHAR(11) NOT NULL,
		UserID INT NOT NULL
	);`

	_, err := db.Exec(sqlStatement)
	if err != nil {
		log.Fatalln(http.StatusInternalServerError)
		log.Fatalln("Error creating database table:", err)
	}

	return err
}