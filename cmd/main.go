package main

import "github.com/gofiber/fiber/v3"

func main() {
	app := fiber.New()
	app.Get("/", func(ctx fiber.Ctx) error {
		return ctx.SendString("hello")
	})
	app.Listen(":9000")

}
