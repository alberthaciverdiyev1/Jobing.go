package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func HomeIndex(c *fiber.Ctx) error {
	return c.Render("main", fiber.Map{
		"title": "Ana Sayfa",
	})
}
