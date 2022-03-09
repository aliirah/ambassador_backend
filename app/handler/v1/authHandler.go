package v1

import (
	"alirah/app/domain"
	"alirah/app/request/v1/auth"
	userResource "alirah/app/resource/user"
	"alirah/database"
	authHelper "alirah/util/auth"
	"alirah/util/rest"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
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

	password, _ := bcrypt.GenerateFromPassword([]byte(body.Password), 12)
	user := domain.User{
		FirstName:    body.FirstName,
		LastName:     body.LastName,
		Email:        body.Email,
		Password:     password,
		IsAmbassador: false,
	}
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

	token, Terr := authHelper.CreateToken(c, user.Id)
	if Terr != nil {
		return rest.BadRequest(c, err)
	}

	return rest.Ok(c, fiber.Map{
		"message": "Login Successfully",
		"user":    userResource.SingleResource(&user),
		"token":   token,
	})
}
