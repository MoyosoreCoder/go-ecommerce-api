// package handler
package handler

import (
	"encoding/json"
	"net/http"

	"github.com/MoyosoreCoder/go-ecommerce-api/internal/registration"
)


// RegisterUser godoc
// @Summary Register a new user
// @Description Creates a new user and hashes the password
// @Tags Users
// @Accept json
// @Produce json
// @Param user body RegisterRequest true "User Registration data"
// @Success 201 {object} SuccessResponse 
// @Failure 400 {object} ErrorResponse
//@Failure 409 {object} ErrorResponse
// @Router /register [post]

type RegisterRequest struct {
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type SuccessResponse struct {
	Message string `json:"message" example:"User registered successfully"`
}

type ErrorResponse struct {
	Error string `json:"error" example:"User already exists"`
}

func RegisterUserHandler (w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost{
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "Invalid request body"})
		return
	}
	if err := registration.RegisterUser(req.Username, req.Email, req.Password); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(SuccessResponse{Message: "User registered successfully"})
}

	
//  Struct for login
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
// LoginUserHandler godoc
// @Summary Login a user
// @Description Authenticates a user with email and password
// @Tags Users
// @Accept json
// @Produce json
// @Param user body LoginRequest true "Login data"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Router /login [post]
func LoginUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "Invalid request body"})
		return
	}

	if err := registration.LoginUser(req.Email, req.Password); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(SuccessResponse{Message: "Login successful"})
}