package main

import (
	"app/matcher"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/django/v3"
)

func main() {
	matcher := matcher.New()
	engine := django.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/", func(c *fiber.Ctx) error {

		return c.Render("start", fiber.Map{
			"gamecode": matcher.getCode(c.IP()),
		})
	})

	log.Fatal(app.Listen(":8080"))
}
