package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/go-sql-driver/mysql"
)

// Global variable to hold the DB connection
var DB *sql.DB

// InitDB initializes the database connection
func InitDB() {
	// Debug: Check CWD
	cwd, _ := os.Getwd()
    fmt.Println("Current working directory:", cwd)

	// Load environment variables from .env file
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	} else {
		fmt.Println(".env file loaded successfully")
	}

	// Get database credentials from environment variables
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// Debug: Test if variables are loaded
    fmt.Println("DB_USER:", dbUser)
	fmt.Println("DB_PASSWORD:", dbPassword)
	fmt.Println("DB_HOST:", dbHost)
	fmt.Println("DB_PORT:", dbPort)
	fmt.Println("DB_NAME:", dbName)

	// Ensure all necessary environment variables are set
	if dbUser == "" || dbPassword == "" || dbHost == "" || dbPort == "" || dbName == "" {
		log.Fatal("Missing environment variables for database connection")
	}

	// MySQL data source name (DSN)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
	// MySQL database credentials
	// dsn := "root:yourpassword@tcp(localhost:3306)/yourdb"
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	// Check if the connection is successful
	if err := DB.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MySQL!")
}
