package models

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client
var db *mongo.Database

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	dbName := os.Getenv("db_name")

	db = client.Database(dbName)
}

// GetDB instance
func GetDB() *mongo.Database {
	return db
}
