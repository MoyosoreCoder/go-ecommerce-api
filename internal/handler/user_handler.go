// package handler
package handler

import (
	"encoding/json"
	"net/http"

	"github.com/MoyosoreCoder/go-ecommerce-api/internal/registration"
)

type RegisterRequest struct {
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"`
}

func RegisterUserHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req RegisterRequest
	// Decode req body
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Call registration logic
	err = registration.RegisterUser(req.Username, req.Email, req.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// On success write response
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("This user is registered successfully"))
}
//  Struct for login
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
func LoginUserHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req LoginRequest

	// Decode request body 
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Call login logic
	err = registration.LoginUser(req.Email, req.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Login successful"))
}