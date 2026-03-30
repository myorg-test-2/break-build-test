package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os/exec"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// SQL Injection
	id := r.URL.Query().Get("id")
	db, _ := sql.Open("mysql", "user:pass@/db")
	db.Query("SELECT * FROM users WHERE id = " + id)

	// Command Injection
	host := r.URL.Query().Get("host")
	out, _ := exec.Command("ping", "-c", "4", host).Output()
	fmt.Fprintf(w, string(out))

	// XSS
	name := r.URL.Query().Get("name")
	fmt.Fprintf(w, "<h1>Hello "+name+"</h1>")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
