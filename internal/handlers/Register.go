package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/MoyosoreCoder/go-ecommerce-api/internal/models"
	"github.com/MoyosoreCoder/go-ecommerce-api/internal/repository"
	"github.com/jackc/pgx/v5/pgxpool"
)

func RegisterUserHandler(pool *pgxpool.Pool) http.HandlerFunc {

// This will return hanndler function for http stuff
	return func(w http.ResponseWriter, r *http.Request) {

		// Create the user in a var variable first
		var user models.User

		// Decode from incoming request
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return 
		}
		err = repository.CreateUser(r.Context(), pool, user)
		if err != nil {
			http.Error(w, "failed to create user", http.StatusInternalServerError)
			return
		}
		//Send a success message
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		json.NewEncoder(w).Encode(map[string] string {
			"message": "User Created Successfully",
		})
	}
}