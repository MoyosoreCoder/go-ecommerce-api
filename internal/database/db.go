// connecting database import package database first
package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect(){
	//get conncetion string from environment variables
	connStr := os.Getenv("DB_URL")
	if connStr == "" {
		log.Fatal("DB_URL environment variable not set")
	}
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil{
		log.Fatal("Error opening database:", err)

	}
	// verify error
	err = DB.Ping()
	if err != nil {
		log.Fatal("Error connecting to database", err)
	}
	log.Println("Database connected successfully")
}