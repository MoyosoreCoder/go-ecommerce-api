	// import package database
	package database

	import (
		"database/sql"
		"log"
		"os"
		_ "github.com/lib/pq"
		"github.com/joho/godotenv"
		"fmt"
	)
	// Global variable holding the database connection pool
	//it allows this DB to be declared at top level.
	var DB *sql.DB

	func Connect(){
		//Load env variables
		err := godotenv.Load()

		if err != nil {
			log.Println("No .env file found, using system env variables")
		}

		// Read environment variables
		host := os.Getenv("DB_HOST")
		port := os.Getenv("DB_PORT")
		user := os.Getenv("DB_USER")
		password := os.Getenv("DB_PASSWORD")
		dbname := os.Getenv("DB_NAME")

		// Check that all variables exist
		if host == "" || port == "" || user == "" || password == "" || dbname == "" {
			log.Fatal("One or more DB environment variables not set")
		}
		log.Println("DB_NAME =", dbname)
		log.Println("DB_HOST =", host)
		log.Println("DB_PORT =", port)
		
		// Build connection string
		connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname,)
		log.Println("CONNECTED TO DATABASE:", connStr)
		// Open connection String
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