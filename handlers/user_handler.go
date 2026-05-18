package handlers

import (
	"encoding/json"
	"httpserver/models"
	"log"
	"net/http"
	"httpserver/utils"

	"golang.org/x/crypto/bcrypt"
	"httpserver/database"
)
//var db = &database.DB


func RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	var userModel models.RegisterUser
	//if all fields are correct, register to database
	if r.Method != http.MethodPost {
		http.Error(w, "request method not acceptable", http.StatusMethodNotAllowed)
		return
	}

	//decode  the body request
	err := json.NewDecoder(r.Body).Decode(&userModel)
	if err!= nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	if userModel.Username == "" || userModel.Email == "" ||userModel.Password == "" {
		http.Error(w, "missing field", http.StatusBadRequest)
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userModel.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error generating hashed password",  http.StatusInternalServerError)
		return
	}
	userModel.Password = string(hashedPassword)
	//save to database
	_, err = database.DB.Exec(
		"INSERT INTO users (username, email, password) VALUES ($1, $2, $3)",
		userModel.Username,
		userModel.Email,
		userModel.Password,
	)
	if err != nil {
		log.Printf("Error: %v", err)
		http.Error(w, "Error saving to database", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "user registered successfully"})

}
//Login Response struct
type LoginResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

//Login handler
func LoginHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

	var requestBody models.LoginUser
    err := json.NewDecoder(r.Body).Decode(&requestBody); 
	if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
	//w.Header().Set("Content-Type", "application/json")
    //json.NewEncoder(w).Encode(map[string]string{"token": "token"})

	// 1. Fetch user from DB based on email (pseudo-code)
	var userID int
	var email string
	var hashedPassword string

	query := "SELECT id, email, password FROM users WHERE email = $1"
	err = database.DB.QueryRow(query, requestBody.Email).Scan(&userID, &email, &hashedPassword)
	
    if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized); 
		return
	}
	// 2. Compare password
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(requestBody.Password))
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}
	// 3. Generate JWT
	token, err := utils.GenerateJWT(userID, email)
	if err != nil {
		http.Error(w, "Could not generate token", http.StatusInternalServerError)
		return
	}

	//login response message
	response := LoginResponse{
		Message: "Login successful",
		Token: token,

	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
func ProfileHandler(w http.ResponseWriter, r *http.Request) {

	userID := r.Context().Value("userID")
	email := r.Context().Value("email")

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]interface{}{
		"user_id": userID,
		"email":   email,
	})
}