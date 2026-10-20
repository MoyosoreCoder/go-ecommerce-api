package models

type GoogleUser  struct {
	ID string `json:"sub"`
	Name string `json:"name"`
	Email string `json:"email"`
}