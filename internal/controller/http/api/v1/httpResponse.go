package v1

import (
	"github.com/gofiber/fiber/v2"
)

func ResponseError(ctf *fiber.Ctx, err error, httpStatusCode int) error {
	return ctf.Status(httpStatusCode).JSON(err.Error())
}

func ResponseOk(ctf *fiber.Ctx, data interface{}) error {
	return ctf.Status(fiber.StatusOK).JSON(data)
}

func ResponseCreated(ctf *fiber.Ctx, data interface{}) error {
	return ctf.Status(fiber.StatusCreated).JSON(data)
}
