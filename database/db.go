package database

import (
	"alirah/config"
	"alirah/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func Connect() {
	var err error
	dsn := config.Dns()

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Could not connect with the database!")
	}
}

func Migrate() {
	err := DB.AutoMigrate(domain.User{})
	if err != nil {
		log.Println("cannot migrate")
	}
}
