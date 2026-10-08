package repositories
import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/MoyosoreCoder/go-ecommerce-api/internal/models"
	"context"
)


//i was thinking of importing database here and model
func CreateUser(pool *pgxpool.Pool, user models.User) {

	query := `INSERT INTO users(name, email, password)
	VALUES ($1, $2, $3)`

	pool.Exec(context.Background(),query, user.Name, user.Email, user.Password)
}