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

// User object that stores a user's properties
type User struct {
	ID       int
	Email    string
	Name     string
	Password string
}

// connect to datatbase
func startDatabase() {
	var psqlInfo string

	// use localhost if environment variable DATABASE_URL is empty
	if _, exist := os.LookupEnv("DATABASE_URL"); exist {
		psqlInfo = os.Getenv("DATABASE_URL")
	} else {
		psqlInfo = fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",
			dbhost, dbport, user, password, dbname)
	}

	var err error
	db, err = sql.Open("postgres", psqlInfo)
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
	mux.HandleFunc("/users", handleUser)

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
