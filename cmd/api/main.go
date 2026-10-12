package main

import (
	"fmt"
	"github.com/MoyosoreCoder/go-ecommerce-api/internal/database"
	"github.com/MoyosoreCoder/go-ecommerce-api/internal/handlers"
	"log"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, welcome to the home handler")
}

func main() {
	pool, err := database.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	http.HandleFunc("/", homeHandler)

	http.HandleFunc("/auth/google", handlers.GoogleLoginHandler)
	http.HandleFunc("/auth/google/callback", handlers.GoogleCallbackHandler)
	http.HandleFunc("/api/register", handlers.RegisterUserHandler(pool))

	fmt.Println("Server is running at http://localhost:8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
