package database

import (
	"log"
	"whalebone-assignment/internal/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Service struct {
	Db *gorm.DB
}

// New creates db connection and runs autoMigrate()
func New(dbName string) Service {
	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&models.DatabaseUser{})

	return Service{
		Db: db,
	}
}
