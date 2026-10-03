 package database


import (
	"context"
	"fmt"
	"os"
	"github.com/joho/godotenv"
	"github.com/jackc/pgx/v5/pgxpool"
)

 func ConnectDB() (*pgxpool.Pool, error) {

	err := godotenv.Load()

	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %w", err)
	}

	pool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}
	
	fmt.Println("Database connected successfully")
	return pool, nil
 }