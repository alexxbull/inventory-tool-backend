package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

// register or sign in in user
func handleUser(w http.ResponseWriter, req *http.Request) {
	// check if request Method is valid
	if req.Method != "POST" {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	// store the request data in a struct
	user := User{}
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		http.Error(w, "incomplete input data", 400)
		return
	}

	// if user's name is empty sign in, else register
	if user.Name == "" {
		signIn(user, w)
	} else {
		register(user, w)
	}
}

func getUser(u User) (User, error) {
	// check if the user's email has already been registered
	sqlGetUser := `
		SELECT * 
		FROM users
		WHERE email = $1`

	dbUser := User{}
	err := db.QueryRow(sqlGetUser, u.Email).Scan(&dbUser.ID, &dbUser.Email, &dbUser.Name, &dbUser.Password)
	return dbUser, err
}

// add user to the database
func register(u User, w http.ResponseWriter) {
	// check if user is already been registered
	_, err := getUser(u)
	if err != sql.ErrNoRows {
		http.Error(w, "email is already registered", 400)
		return
	}

	// add user to the database
	sqlInsert := `
		INSERT INTO users (Email, Name, Password)
		VALUES ($1, $2, $3);`

	// get hashed password
	hashedPassword, err := hashPassword(u.Password)
	if err != nil {
		http.Error(w, "error hashing password", 500)
	}

	_, err = db.Exec(sqlInsert, u.Email, u.Name, hashedPassword)
	if err != nil {
		http.Error(w, "unable to add user", 500)
		return
	}

	// get newly created user from database
	dbUser, err := getUser(u)
	if err != nil {
		http.Error(w, "user is not registered", 404)
		return
	}

	// respond with new user's ID
	sendUserID(dbUser, w)
}

// validate user information
func signIn(u User, w http.ResponseWriter) {
	// retrieve the user information from the database
	dbUser, err := getUser(u)
	if err != nil {
		http.Error(w, "user is not registered", 404)
		return
	}

	// check if the user's password matches the hashed password in the database
	passwordError := verifyPassword(u.Password, dbUser.Password)
	if !passwordError {
		http.Error(w, "invalid credentials", 400)
		return
	}

	// respond with user ID
	sendUserID(dbUser, w)
}

// send a HTTP resposne with the given user's ID
func sendUserID(u User, w http.ResponseWriter) {
	jsonID, err := json.Marshal(struct{ UserID int }{u.ID})
	if err != nil {
		http.Error(w, "unable to return user id", 500)
		return
	}
	w.Write(jsonID)
}

// return a hash generated from the given string
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func verifyPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
