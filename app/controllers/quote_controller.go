package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func GetQuotes(c *fiber.Ctx) error {

	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		"quotes": "return data here",
	})
}