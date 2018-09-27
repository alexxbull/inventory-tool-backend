package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
	"github.com/rs/cors"
)

const (
	// database constants
	dbhost   = "localhost"
	dbport   = 5432
	user     = "alexx"
	password = "password"
	dbname   = "inventory"
)

var (
	db   *sql.DB
	mux  *http.ServeMux
	port string
)

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
	var err error
	db, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	fmt.Println(os.Getenv("DATABASE_URL"))
	fmt.Println(err)
	if err != nil {
		log.Fatal(err)
	}
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

	// startup server message to show running successfully
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Server running on port", port)
	})
}

func main() {
	// connect to database and build endpoints
	startDatabase()
	handleRoutes()

	// close database connection
	defer db.Close()

	// enable cors
	handler := cors.AllowAll().Handler(mux)

	// get server port number
	port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// start server
	addr := fmt.Sprintf(":%s", port)
	http.ListenAndServe(addr, handler)
}
