package auth

import (
	"alirah/app/domain"
	"alirah/app/middleware"
	"alirah/app/request/v1/auth"
	userResource "alirah/app/resource/user"
	"alirah/database"
	authHelper "alirah/util/auth"
	"alirah/util/rest"
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
		"message":    "Hello World!",
		"ambassador": userResource.SingleResource(&user),
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
		return rest.BadRequest(c, Terr)
	}

	return rest.Ok(c, fiber.Map{
		"message":    "Login Successfully",
		"ambassador": userResource.SingleResource(&user),
		"token":      token,
	})
}

func User(c *fiber.Ctx) error {
	user := middleware.GetUser(c)

	return rest.Ok(c, fiber.Map{
		"message":    "Successfully",
		"ambassador": userResource.SingleResource(&user),
	})
}

func Logout(c *fiber.Ctx) error {
	authHelper.RemoveAuthCookie(c)

	return rest.Ok(c, fiber.Map{
		"message": "Successfully Logout",
	})
}

func UpdateInfo(c *fiber.Ctx) error {
	var body auth.UpdateUserData

	if err := c.BodyParser(&body); err != nil {
		return rest.BadRequest(c, err)
	}

	err := auth.UpdateUserValidate(&body)
	if err != nil {
		return rest.ValidationError(c, err)
	}

	user := middleware.GetUser(c)

	var existsUser domain.User
	res := database.DB.Find(&existsUser, "email = ?", body.Email)

	if res.RowsAffected > 0 && existsUser.Id != user.Id {
		return rest.BadRequest(c, "email is taken")
	}

	database.DB.Model(&user).Updates(domain.User{
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Email:     body.Email,
	})

	return rest.Ok(c, fiber.Map{
		"message":    "ambassador " + user.Email + " successfully updated.",
		"ambassador": userResource.SingleResource(&user),
	})
}

func UpdatePassword(c *fiber.Ctx) error {
	var body auth.UpdatePasswordData

	if err := c.BodyParser(&body); err != nil {
		return rest.BadRequest(c, err)
	}

	err := auth.UpdatePasswordValidate(&body)
	if err != nil {
		return rest.ValidationError(c, err)
	}

	user := middleware.GetUser(c)
	user.SetPassword(body.Password)

	database.DB.Model(&user).Updates(&user)
	authHelper.RemoveAuthCookie(c)
	return rest.Ok(c, fiber.Map{
		"message": "password of ambassador " + user.Email + " successfully updated.",
	})
}
