// D:\Projects\REST-API-Go\pkg\database\connection.go
package database

import (
	"fmt"
	"os"

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
	db, err = gorm.Open("mysql", connectionString)

	if err != nil {
		fmt.Println("Error opening database:", err)
		return
	}

	// Test the connection.
	err = db.DB().Ping()
	if err != nil {
		fmt.Println("Error pinging database:", err)
		return
	}

	fmt.Println("Database 🗃️  connected successfully.")
}

// GetDB returns the database instance.
func GetDB() *gorm.DB {
	return db
}
