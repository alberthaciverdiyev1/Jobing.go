package main

import (
	"Jobing/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/jet/v2"
)

func main() {
	engine := jet.New("./views", ".jet")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	routes.WebRoutes(app)

	err := app.Listen(":3000")

	if err != nil {
		return
	}
}
