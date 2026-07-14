package main

import (
	"fmt"
	"net/http"
	"log"
	"github.com/MoyosoreCoder/go-ecommerce-api/internal/database"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, welcome to the home handler")
}

func main() {
	database.ConnectDB()
	http.HandleFunc("/", homeHandler)
	fmt.Println("Server is running at http://localhost:8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}