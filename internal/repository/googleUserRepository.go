package repository





import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/MoyosoreCoder/go-ecommerce-api/internal/models"

)

func SaveGoogleUser(ctx context.Context, pool *pgxpool.Pool, user models.GoogleUser) {

	// query string
	query := `
	INSERT into users(google_id, name, email) VALUES ($1, $2, $3) RETURNING id
	`
	// the id returned can be used for jwt for authentication process, later on
	err := pool.QueryRow(ctx, query, user.ID, user.Name, user.Email)

}
