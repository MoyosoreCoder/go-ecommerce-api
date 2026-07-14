package database

import (
	"fmt"
	"log"
	"os"
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func ConnectDB() (*pgxpool.Pool, error) {
	// 1. The first thing we need to do is load our environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file or missing")
	}
	
	// 2. We need the config from the official docs
	config, err := pgxpool.ParseConfig(os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Println("Unable to parse configuration string")
		log.Fatal(err)
	}

	//3. config after connect
	config.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
    // do something with every new connection--set time zone
	_, err = conn.Exec(ctx, "set timezone to 'UTC';")
	return err
	}
	
	// 4. Now we can create the pool
	
	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		fmt.Println("Unable to create connection pool")
		log.Fatal(err)
	}
	fmt.Println("Successfully connected to the database")
	return pool, nil	
}