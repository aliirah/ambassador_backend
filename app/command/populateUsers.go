package main

import (
	"alirah/app/domain"
	"alirah/database"
	"github.com/bxcodec/faker/v3"
	"math/rand"
	"time"
)

func main() {
	database.Connect()
	for i := 0; i < 30; i++ {
		ambassador := domain.User{
			FirstName:    faker.FirstName(),
			LastName:     faker.LastName(),
			Email:        faker.Email(),
			IsAmbassador: RandBool(),
		}
		ambassador.SetPassword("password")

		database.DB.Create(&ambassador)
	}
}

func RandBool() bool {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(2) == 1
}
