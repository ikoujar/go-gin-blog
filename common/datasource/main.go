package datasource

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func ConnectDatabase(dsn string) {
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("Database connection failed")
		return
	}
	log.Println("Database connected")
	DB = database
}

func AutoMigrate(dst interface{}) {
	err := DB.AutoMigrate(dst)
	if err != nil {
		return
	}
}
