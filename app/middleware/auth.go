package middleware

import (
	"alirah/app/domain"
	"alirah/database"
	authHelper "alirah/util/auth"
	"alirah/util/rest"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func IsAuth(c *fiber.Ctx) error {
	jwtCookie := c.Cookies("jwt")
	_, err := authHelper.ParseToken(jwtCookie)
	if err != nil {
		return rest.Unauthorized(c)
	}
	return c.Next()
}

func IsGuest(c *fiber.Ctx) error {
	jwtCookie := c.Cookies("jwt")
	_, err := authHelper.ParseToken(jwtCookie)
	if err != nil {
		return c.Next()
	}
	return rest.BadRequest(c, "Already Login")
}

func GetUserId(c *fiber.Ctx) (uint, error) {
	jwtCookie := c.Cookies("jwt")
	token, err := authHelper.ParseToken(jwtCookie)
	if err != nil {
		return 0, rest.Unauthorized(c)
	}

	payload := token.Claims.(*jwt.StandardClaims)

	id, _ := strconv.Atoi(payload.Subject)

	return uint(id), nil
}

func GetUser(c *fiber.Ctx) domain.User {
	var user domain.User

	id, _ := GetUserId(c)
	database.DB.Where("id = ?", id).First(&user)

	return user
}
