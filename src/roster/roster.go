package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"roster/web"
	// "net/smtp"
	"os"
)

var port string

func init() {
	_, err := sql.Open("postgres", "user=pqgotest dbname=pqgotest sslmode=verify-full")
	if err != nil {
		log.Fatal(err)
	}

	port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
}

func main() {
	err := http.ListenAndServe(":"+port, web.Router())
	panic(err)
}
