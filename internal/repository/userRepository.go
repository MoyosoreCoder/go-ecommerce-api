package repositories
import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/MoyosoreCoder/go-ecommerce-api/internal/models"
	"context"
)


//i was thinking of importing database here and model
func CreateUser(ctx context.Context, pool *pgxpool.Pool, user models.User) error {

	query := `INSERT INTO users(name, email, password)
	VALUES ($1, $2, $3)`

	_, err := pool.Exec(ctx,query, user.Name, user.Email, user.Password)
	if err != nil {
		return err
	}
	return  nil
}