// package registration
package registration
import (
	"errors"
	"github.com/MoyosoreCoder/go-ecommerce-api/internal/database"
	"golang.org/x/crypto/bcrypt"
)
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
