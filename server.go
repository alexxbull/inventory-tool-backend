package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/rs/cors"

	_ "github.com/lib/pq"
)

const (
	port = 8080
	// database constants
	dbhost   = "localhost"
	dbport   = 5432
	user     = "alexx"
	password = "password"
	dbname   = "inventory"
)

var db *sql.DB
var mux *http.ServeMux

func init() {
	startDatabase()
	handleRoutes()
}

// Item object that stores an item's properties
type Item struct {
	ID        int
	Name      string
	Price     float64
	Color     string
	Condition string
	UserID    int
}

// connect to datatbase
func startDatabase() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		dbhost, dbport, user, password, dbname)

	var err error
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	// check if connected to database
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to database.")
}

// initialize valid server routes
func handleRoutes() {
	mux = http.NewServeMux()

	// DELETE
	mux.HandleFunc("/items/delete", deleteItem)

	// GET
	mux.HandleFunc("/items/view", viewItems)

	// POST
	mux.HandleFunc("/items/add", addItem)

	// PUT
	mux.HandleFunc("/items/edit", editItem)

	mux.HandleFunc("/", startServer)
}

func startServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Server now running...")
}

func main() {
	// close database connection
	defer db.Close()

	// enable cors
	handler := cors.AllowAll().Handler(mux)

	// start server
	addr := fmt.Sprintf(":%d", 8080)
	http.ListenAndServe(addr, handler)
}
