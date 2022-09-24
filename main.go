package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

// may need to wrap html with template.HTML() when sent to vue to replace the existing contents of an element
func setup() error {
	app := fiber.New(fiber.Config{
		Views: html.New("templates", ".html"),
	})

	// initially rendered html
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			// could pass an initial list of animals here
		})
	})

	// send an updated list of animals in HTML format, then petite vue will update the contents of an HTML element
	app.Get("/animals", func(c *fiber.Ctx) error {
		return c.Render("animals", fiber.Map{
			"Animals": []string{"Percy", "Friday", "Shorty", "Frost"},
		})
	})

	return app.Listen(":3000")
}

func main() {
	if err := setup(); err != nil {
		fmt.Println("failed to setup app: %w", err)
	}
}
