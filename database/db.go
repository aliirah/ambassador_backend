package database

import (
	"alirah/app/domain/product"
	userDomain "alirah/app/domain/user"
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"strconv"
)

var DB *gorm.DB

func Connect() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var (
		host     = os.Getenv("POSTGRES_HOST")
		port, _  = strconv.Atoi(os.Getenv("POSTGRES_PORT"))
		user     = os.Getenv("POSTGRES_USER")
		password = os.Getenv("POSTGRES_PASSWORD")
		dbname   = os.Getenv("POSTGRES_DB")
	)

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Could not connect with the database!")
	}
}

func Migrate() {
	err := DB.AutoMigrate(
		userDomain.User{},
		product.Product{},
	)
	if err != nil {
		log.Println("cannot migrate")
	}
}
