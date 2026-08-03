package database

import (
	"fmt"
	"os"
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func ConnectDB() (*pgxpool.Pool, error) {
	// 1. The first thing we need to do is load our environment variables
	// with Errorf, you can easily use format specifier
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("Error loading .env file: %w", err)
	}
	
	// 2. Next is the onfig variable from the official docs
	config, err := pgxpool.ParseConfig(os.Getenv("DATABASE_URL"))
	if err != nil {
		return  nil, fmt.Errorf("Unable to parse database config: %w", err)
	}

	//3. config after connect
	config.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
    // do something with every new connection--set time zone
	_, err := conn.Exec(ctx, "set timezone to 'UTC';")
	return err
	}
	
	// 4. Now we can create the pool
	
	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, fmt.Errorf("Unable to create connection pool: %w", err)
	}
	fmt.Println("Successfully connected to the database")
	return pool, nil
	
}