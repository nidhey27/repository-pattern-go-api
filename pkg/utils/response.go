package utils

import "github.com/gofiber/fiber/v2"

func SendResponse(statusCode int, message string, err string, data interface{}, ctx *fiber.Ctx) error {
	return ctx.Status(statusCode).JSON(&fiber.Map{
		"code":    statusCode,
		"message": message,
		"error":   err,
		"data":    data,
	})
}
