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
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, payload).SignedString([]byte(secretKey))

	if err != nil {
		return "", err
	}

	AddAuthCookie(c, token, expireTime)

	return token, nil
}

func AddAuthCookie(c *fiber.Ctx, token string, expireTime time.Time) {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  expireTime,
		HTTPOnly: true,
	}
	c.Cookie(&cookie)
}

func RemoveAuthCookie(c *fiber.Ctx) {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)
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
