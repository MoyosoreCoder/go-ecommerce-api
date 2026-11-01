package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/MoyosoreCoder/go-ecommerce-api/internal/database"
	"github.com/MoyosoreCoder/go-ecommerce-api/internal/handlers"
	"github.com/joho/godotenv"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, welcome to the home handler")
}

func main() {
	// load env variables for database connection
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/auth/google", handlers.GoogleLoginHandler)
	http.HandleFunc("/auth/google/callback", handlers.GoogleCallbackHandler)

	pool, err := database.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	fmt.Println("Server is running at http://localhost:8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
