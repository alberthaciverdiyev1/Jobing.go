package routes

import (
	"Jobing/controllers"
	"github.com/gofiber/fiber/v2"
)

func WebRoutes(app *fiber.App) {
	app.Get("/", controllers.HomeIndex)
}
