package main

import (
	"log"

	"github.com/creative201347/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})
	app.Post("/api/products", handlers.CreateProduct)
	app.Get("/api/products", handlers.GetAllProducts)

	log.Fatal(app.Listen(":3000"))
}
