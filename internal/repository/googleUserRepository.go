package repository

import (
	"context"

	"github.com/MoyosoreCoder/go-ecommerce-api/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

func SaveGoogleUser(ctx context.Context, pool *pgxpool.Pool, user models.GoogleUser) (int, error) {

	// query string
	// the id returned can be used for jwt for authentication process, later on

	query := `
	
	INSERT into users(google_id, name, email) VALUES ($1, $2, $3) RETURNING id
	`

	var userID int

	err := pool.QueryRow(ctx, query, user.ID, user.Name, user.Email).Scan(&userID)
	if err != nil {
		return 0, err

	}
	return userID, nil
	

}

