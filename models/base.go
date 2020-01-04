package models

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var db *gorm.DB

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	dbHost := os.Getenv("db_host")
	dbPort := os.Getenv("db_port")
	dbName := os.Getenv("db_name")
	dbUser := os.Getenv("db_user")
	dbPassword := os.Getenv("db_pass")

	dbURI := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable", dbHost, dbPort, dbName, dbUser, dbPassword)
	log.Println("Connecting to " + dbURI)

	conn, err := gorm.Open("postgres", dbURI)
	if err != nil {
		log.Fatal(err)
	}

	db = conn
	db.Debug().AutoMigrate(
		&User{},
		&Product{},
	)
}

// GetDB instance
func GetDB() *gorm.DB {
	return db
}
