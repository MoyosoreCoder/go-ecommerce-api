package handlers

import (
	"encoding/json"

	"log"
	"net/http"

	"github.com/MoyosoreCoder/go-ecommerce-api/models"
	"github.com/MoyosoreCoder/go-ecommerce-api/utils"

	"github.com/MoyosoreCoder/go-ecommerce-api/database"
	"golang.org/x/crypto/bcrypt"
)

type Response struct {
    Message string      `json:"message,omitempty"`
    Error   string      `json:"error,omitempty"`
    Data    interface{} `json:"data,omitempty"`
}

func RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
        var userModel models.RegisterUser
        //if all fields are correct, register to database
        if r.Method != http.MethodPost {
                
                w.Header().Set("Content-Type", "application/json")
                w.WriteHeader(http.StatusMethodNotAllowed)
                json.NewEncoder(w).Encode(Response{
                        Error: "method not allowed",
                })
                return
        }

        //decode  the body request
        err := json.NewDecoder(r.Body).Decode(&userModel)
        if err!= nil {
                w.Header().Set("Content-Type", "application/json")
                w.WriteHeader(http.StatusBadRequest)
                json.NewEncoder(w).Encode(Response{
                        Error: "invalid request body",
                })
                return
        }
        var missingFields []string
        if userModel.Username == "" {
                missingFields = append(missingFields, "username")
        }
        if userModel.Email == "" {
                missingFields = append(missingFields, "email")
        }
        if userModel.Password == "" {
                missingFields = append(missingFields, "password")
        }
        if len(missingFields) > 0 {
                w.Header().Set("Content-Type", "application/json")
                w.WriteHeader(http.StatusBadRequest)
                json.NewEncoder(w).Encode(Response{
                        Error: "missing required fields",
                        Data:  missingFields,
                })
                return
        }
        hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userModel.Password), bcrypt.DefaultCost)
        if err != nil {
                w.Header().Set("Content-Type", "application/json")
                w.WriteHeader(http.StatusInternalServerError)
                json.NewEncoder(w).Encode(Response{
                        Error: "error hashing password",
                })
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
                w.Header().Set("Content-Type", "application/json")
                 w.WriteHeader(http.StatusInternalServerError)
                 json.NewEncoder(w).Encode(Response{
                        Error: "error creating user",
                })
                return
        }
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusCreated)
        json.NewEncoder(w).Encode(Response{
                Message: "user created successfully",
                Data:   map[string] string {
                        "email": userModel.Email, 
                        "username": userModel.Username,
                },
        })
}


//Login handler
func LoginHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusMethodNotAllowed)
        json.NewEncoder(w).Encode(Response{
                Error: "method not allowed",
        })
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
                w.Header().Set("Content-Type", "application/json")
                 w.WriteHeader(http.StatusNotFound)
                json.NewEncoder(w).Encode(Response{
                        Error: "invalid credentials",
                })
                return
        }
        // 2. Compare password
        err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(requestBody.Password))
        if err != nil {
                w.Header().Set("Content-Type", "application/json")
                w.WriteHeader(http.StatusUnauthorized)
                json.NewEncoder(w).Encode(Response{
                        Error: "invalid credentials",
                })
        }
        // 3. Generate JWT
        token, err := utils.GenerateJWT(userID, email)
        if err != nil {
               w.Header().Set("Content-Type", "application/json")
               w.WriteHeader(http.StatusInternalServerError)
               json.NewEncoder(w).Encode(Response{
                        Error: "error generating token",
                })
                return
        }

       
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        json.NewEncoder(w).Encode(Response{
                Message: "login successful",
                Data: map[string] interface{}{
                        "token": token,
                        "email": email,
                        "id": userID,
                        
                },
        })
}


func ProfileHandler(w http.ResponseWriter, r *http.Request) {

        userID := r.Context().Value("userID")
        email := r.Context().Value("email")

        w.Header().Set("Content-Type", "application/json")

        json.NewEncoder(w).Encode(Response{
                Data: map[string]interface{}{
                        "id": userID,
                        "email": email,
                },
        })
}