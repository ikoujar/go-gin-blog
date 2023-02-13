package config

import (
	"github.com/joho/godotenv"
	"go-gin-course/common/datasource"
	"log"
	"os"
)

// var DB *mongo.Client

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	log.Println("Environment variables initialization is done.")

	dsn := os.Getenv("DB_URI")
	datasource.ConnectDatabase(dsn)

}

//func GetCollection(collectionName string) *mongo.Collection {
//	collection := DB.Database("golang").Collection(collectionName)
//	return collection
//}
