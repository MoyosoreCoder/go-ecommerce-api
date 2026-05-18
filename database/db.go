package database

import (
	"database/sql"
	"fmt"
	"log"
	_"github.com/lib/pq"
	"os"
)


 var DB *sql.DB
 
func ConnectDatabase() {
	
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	// Create the connection string
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	host, port, user, password, dbname,)

	var errDB  error
	// Open the connection
	DB, errDB = sql.Open("postgres", psqlInfo)
	if errDB != nil {
		log.Fatalf("Error opening database: %v", errDB)
	}

	// Verify the connection
	errDB = DB.Ping()
	if errDB != nil {
		log.Fatalf("Error connecting to database: %v", errDB)
	}
	//database connect Successfully if no error
	fmt.Println("Wow, Successfully connected to the database !")	
}