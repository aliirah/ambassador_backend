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

	ambassadorSeeder()
	productSeeder()
	orderSeeder()
}

func ambassadorSeeder() {
	for i := 0; i < 30; i++ {
		ambassador := domain.User{
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
		product := domain.Product{
			Title:       faker.Word(),
			Description: faker.Sentence(),
			Image:       "https://images.unsplash.com/photo-1453728013993-6d66e9c9123a?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxzZWFyY2h8Mnx8dmlld3xlbnwwfHwwfHw%3D&w=1000&q=80",
			Price:       randFloat(),
		}

		database.DB.Create(&product)
	}
}

func orderSeeder() {
	for i := 0; i < 5; i++ {
		var orderItems []domain.OrderItem
		for j := 0; j < rand.Intn(5); j++ {
			price := float64(rand.Intn(90) + 10)
			quantity := uint(rand.Intn(5))

			orderItems = append(orderItems, domain.OrderItem{
				ProductTitle:      faker.Word(),
				Price:             price,
				Quantity:          quantity,
				AdminRevenue:      0.9 + price*float64(quantity),
				AmbassadorRevenue: 0.1 + price*float64(quantity),
			})
		}

		order := domain.Order{
			UserId:          uint(rand.Intn(30) + 1),
			Code:            faker.Username(),
			AmbassadorEmail: faker.Email(),
			FirstName:       faker.FirstName(),
			LastName:        faker.LastNameTag,
			Email:           faker.Email(),
			Complete:        randBool(),
			OrderItem:       orderItems,
		}

		database.DB.Create(&order)
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
