package controllers

import "github.com/gofiber/fiber/v2"

func ContactUs(c *fiber.Ctx) error {
	return c.Render("pages/contact/index", fiber.Map{
		"title": "Contactus",
	})
}

func AboutUs(c *fiber.Ctx) error {
	return c.Render("pages/about/index", fiber.Map{
		"title": "About",
	})
}
