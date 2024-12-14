package database

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

// Global variable to hold the DB connection
var DB *sql.DB

// InitDB initializes the database connection
func InitDB() {
	var err error
	// MySQL database credentials
	dsn := "root:yourpassword@tcp(localhost:3306)/yourtable"
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
