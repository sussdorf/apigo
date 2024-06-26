package controller

import (
	"api/model"

	"github.com/gofiber/fiber/v2"
)

func Token(c *fiber.Ctx) error {
	var input model.AuthenticationInput
	if err := c.BodyParser(&input); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := model.Validate.Struct(&input); err != nil {
		return err
	}

	user, err := model.FindUserByUsername(input.Username)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "User was not Found")
	}

	if err := user.ValidatePassword(input.Password); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Wrong Credentials")
	}

	apiKey, err := model.FindUserKey(user.ID)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"token": apiKey.Token,
		"success": true,
	})
}
