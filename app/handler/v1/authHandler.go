package v1

import (
	"alirah/app/domain"
	"alirah/app/request/v1/auth"
	userResource "alirah/app/resource/user"
	"alirah/database"
	authHelper "alirah/util/auth"
	"alirah/util/rest"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	var body auth.RegisterData
	if err := c.BodyParser(&body); err != nil {
		return rest.BadRequest(c, err)
	}

	err := auth.RegisterValidate(&body)
	if err != nil {
		return rest.ValidationError(c, err)
	}

	user := domain.User{
		FirstName:    body.FirstName,
		LastName:     body.LastName,
		Email:        body.Email,
		IsAmbassador: false,
	}
	user.SetPassword(body.Password)

	res := database.DB.Create(&user)
	if res.Error != nil {
		return rest.BadRequest(c, res.Error)
	}

	return rest.Ok(c, fiber.Map{
		"message": "Hello World!",
		"user":    userResource.SingleResource(&user),
	})
}

func Login(c *fiber.Ctx) error {
	var body auth.LoginData
	if err := c.BodyParser(&body); err != nil {
		return rest.BadRequest(c, err)
	}

	err := auth.LoginValidate(&body)
	if err != nil {
		return rest.ValidationError(c, err)
	}

	var user domain.User
	database.DB.
		Where("email = ?", body.Email).
		Find(&user)

	// Todo fix CreateToken
	token, Terr := authHelper.CreateToken(c, user.Id)
	if Terr != nil {
		return rest.BadRequest(c, Terr)
	}

	return rest.Ok(c, fiber.Map{
		"message": "Login Successfully",
		"user":    userResource.SingleResource(&user),
		"token":   token,
	})
}

func User(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	token, err := authHelper.ParseToken(cookie)
	if err != nil {
		return rest.Unauthorized(c)
	}

	payload := token.Claims.(*jwt.StandardClaims)
	var user domain.User
	fmt.Println(payload.Subject)

	return rest.Ok(c, fiber.Map{
		"payload": payload,
	})

	// Todo fix Get Payload

	database.DB.Where("id = ?", payload.Subject).First(&user)

	return rest.Ok(c, fiber.Map{
		"message": "Successfully",
		"user":    userResource.SingleResource(&user),
	})
}
