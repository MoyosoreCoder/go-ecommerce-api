package repository
import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/MoyosoreCoder/go-ecommerce-api/internal/models"
	"context"
)



func CreateUser(ctx context.Context, pool *pgxpool.Pool, user models.User) error {

	query := `INSERT INTO users(name, email, password)
	VALUES ($1, $2, $3)`

	//pool manages database connection between go application and postgresql
	_, err := pool.Exec(ctx, query, user.Name, user.Email, user.Password)
	if err != nil {
		return err
	}
	return  nil
}