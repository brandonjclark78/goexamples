package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func insecureLogin(db *sql.DB, username, password string) bool {
	query := "SELECT COUNT(*) FROM users WHERE username = '" + username + "' AND password = '" + password + "'"
	var count int
	err := db.QueryRow(query).Scan(&count)
	if err != nil {
		log.Println("Database error:", err)
		return false
	}
	return count > 0
}

func loginHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	password := r.URL.Query().Get("password")

	if insecureLogin(db, username, password) {
		fmt.Fprintf(w, "Login successful!")
	} else {
		fmt.Fprintf(w, "Invalid credentials.")
	}
}

func main() {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		loginHandler(db, w, r)
	})

	log.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}
