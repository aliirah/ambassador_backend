package auth

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
	"time"
)

func CreateToken(c *fiber.Ctx, userId uint) (string, error) {
	var err error
	expireTime := time.Now().Add(time.Hour * 24)

	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	secretKey := os.Getenv("SECRET_KEY")

	// Creating Access Token
	payload := jwt.StandardClaims{
		Subject:   strconv.Itoa(int(userId)),
		ExpiresAt: expireTime.Unix(),
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodES256, payload).SignedString([]byte(secretKey))

	if err != nil {
		return "", err
	}

	// Create cookie
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  expireTime,
		HTTPOnly: true,
	}
	c.Cookie(&cookie)

	return token, nil
}

func ParseToken(cookie string) (*jwt.Token, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	secretKey := os.Getenv("SECRET_KEY")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	return token, nil
}
