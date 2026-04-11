// package registration
package registration
import (
	"errors"
	"github.com/MoyosoreCoder/go-ecommerce-api/internal/database"
	"golang.org/x/crypto/bcrypt"
)
// Register user function
func RegisterUser(username, email, password string) error {
	if username == "" || email == "" || password == "" {
		return errors.New("all fields are required")
	}
	//hash password
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return err
	}
	
	// Insert into DB
	_, err = database.DB.Exec(
		"INSERT INTO users (username, email, password) VALUES ($1, $2, $3)",
		username,
		email,
		string(hashedPassword),
	)
	if err != nil {
		return err
	}
	return  nil

}
// Login user function
func LoginUser(email, password string) error {

	if email == "" || password == "" {
		return errors.New("email and password are required")
	}

	var storedHashedPassword string

	// Get hashed password from DB
	err := database.DB.QueryRow(
		"SELECT password FROM users WHERE email = $1",
		email,
	).Scan(&storedHashedPassword)

	if err != nil {
		return errors.New("invalid email or password")
	}

	// Compare passwords
	err = bcrypt.CompareHashAndPassword(
		[]byte(storedHashedPassword),
		[]byte(password),
	)

	if err != nil {
		return errors.New("invalid email or password")
	}

	return nil
}