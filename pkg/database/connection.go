package database

import (
	"fmt"
	"os"
	"rest-api-redis/pkg/models"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func init() {
	// Initialize the database connection here.
	connectionString := "dan:secret@tcp(localhost:3306)/dapi_assg_db?parseTime=true"
	if os.Getenv("DB_HOST") != "" && os.Getenv("DB_PORT") != "" && os.Getenv("DB_USER") != "" && os.Getenv("DB_PASSWORD") != "" && os.Getenv("DB_NAME") != "" {
		DB_HOST := os.Getenv("DB_HOST")
		DB_PORT := os.Getenv("DB_PORT")
		DB_USER := os.Getenv("DB_USER")
		DB_PASSWORD := os.Getenv("DB_PASSWORD")
		DB_NAME := os.Getenv("DB_NAME")
		connectionString = fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true", DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME)
	}

	var err error
	maxRetries := 10
	for retry := 1; retry <= maxRetries; retry++ {
		db, err = gorm.Open("mysql", connectionString)

		if err == nil {
			// Test the connection.
			err = db.DB().Ping()
			if err == nil {
				break // Connection successful, break out of the loop.
			}
		}

		if retry < maxRetries {
			fmt.Printf("Error opening database (Attempt %d of %d): %v\n", retry, maxRetries, err)
			time.Sleep(5 * time.Second) // Wait for 5 seconds before retrying.
		} else {
			fmt.Println("Maximum retries reached. Couldn't establish a database connection.")
			return
		}
	}

	// Run Migrate
	db.Begin().AutoMigrate(&models.User{})

	fmt.Println("Database 🗃️  connected successfully.")
}

// GetDB returns the database instance.
func GetDB() *gorm.DB {
	return db
}
