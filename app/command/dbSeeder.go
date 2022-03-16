package main

import (
	productDomain "alirah/app/domain"
	"alirah/database"
	"github.com/bxcodec/faker/v3"
	"math/rand"
	"time"
)

func main() {
	database.Connect()

	ambassadorSeeder()
	productSeeder()
}

func ambassadorSeeder() {
	for i := 0; i < 30; i++ {
		ambassador := productDomain.User{
			FirstName:    faker.FirstName(),
			LastName:     faker.LastName(),
			Email:        faker.Email(),
			IsAmbassador: randBool(),
		}
		ambassador.SetPassword("password")

		database.DB.Create(&ambassador)
	}
}

func productSeeder() {
	for i := 0; i < 5; i++ {
		product := productDomain.Product{
			Title:       faker.Word(),
			Description: faker.Sentence(),
			Image:       "https://images.unsplash.com/photo-1453728013993-6d66e9c9123a?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxzZWFyY2h8Mnx8dmlld3xlbnwwfHwwfHw%3D&w=1000&q=80",
			Price:       randFloat(),
		}

		database.DB.Create(&product)
	}
}

func randBool() bool {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(2) == 1
}

func randFloat() float64 {
	seed := time.Now().Unix()
	s := rand.NewSource(seed)
	return rand.New(s).Float64()
}
