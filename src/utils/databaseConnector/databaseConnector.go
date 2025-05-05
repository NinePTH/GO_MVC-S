package databaseConnector

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	// _ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

// Global variable to hold the DB connection
var DB *sql.DB

// InitDB initializes the database connection
func InitDB() {
	// Debug: Check CWD
	cwd, _ := os.Getwd()
    fmt.Println("Current working directory:", cwd)

	envPaths := []string{
		"../etc/secrets/.env",    // Local development path
		"/etc/secrets/.env",      // Production path
		"./etc/secrets/.env",     // Alternative local path
	}

	// Load environment variables from different possible locations
	var envLoaded bool
	for _, path := range envPaths {
		absPath, _ := filepath.Abs(path)
		if err := godotenv.Load(path); err == nil {
			fmt.Printf(".env file loaded successfully from: %s\n", absPath)
			envLoaded = true
			break
		}
	}

	if !envLoaded {
		// If no .env file is found, check if we're in production with environment variables already set
		if os.Getenv("DB_USER") != "" {
			fmt.Println("No .env file found, but environment variables are set (production environment)")
		} else {
			log.Fatal("Error: Could not load .env file and no environment variables are set")
		}
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

	// PostgreSQL data source name (DSN)
	// Add require &pool_mode=session before merge to main
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable&pool_mode=session", dbUser, dbPassword, dbHost, dbPort, dbName)
	// MySQL database credentials
	// dsn := "root:yourpassword@tcp(localhost:3306)/yourdb"
	var err error
	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	// Check if the connection is successful
	if err := DB.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to PostgreSQL!")
}
