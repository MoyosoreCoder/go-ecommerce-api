package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func ConnectDB() {
	// 1. loading environment variables
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file or missing")
	}

	// 2. for parsing connection string config in docs
	config, err := pgxpool.ParseConfig(os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Println("Unable to Parse connection string Config")
		log.Fatal(err) 	 		
	}
	// 3. creates new pool with the given config
	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		fmt.Println("Unable to establish a new database connection config pool")
		log.Fatal(err)
	}
	// 4. use pool here to close every connection for safety
	defer pool.Close()

	// 5. ping() checks if the database is reachable and returns an error if not
	err = pool.Ping(context.Background())
	if err != nil {
		fmt.Println("Unable to ping the Database")
		log.Fatal(err)
	}

	// 6. successful connection message
	fmt.Println("Successfully Connected to the Database")
}