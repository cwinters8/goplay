package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/gofiber/websocket/v2"
)

func setup() error {
	app := fiber.New(fiber.Config{
		Views: html.New("templates", ".html"),
	})

	// initially rendered html
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Animals": []string{"Four", "Speck", "Hush"},
		})
	})

	// send an updated list of animals in HTML format, then petite vue will update the contents of an HTML element
	app.Get("/animals", func(c *fiber.Ctx) error {
		return c.Render("animals", fiber.Map{
			"Animals": []string{"Percy", "Friday", "Shorty", "Frost"},
		})
	})

	// websockets!
	ws := app.Group("/ws")
	ws.Use() // TODO: websocket middleware
	ws.Get("/:id", websocket.New(func(c *websocket.Conn) {
		for {
			// mt (message type) could be either binary type or text (I think?) type
			_, _, err := c.ReadMessage()
			if err != nil {
				break
			}
			// do something with msg
			// for now, echo back to the user?
		}
		fmt.Printf("connection closed") // maybe pass the id param here
	}))

	return app.Listen(":3000")
}

func main() {
	if err := setup(); err != nil {
		fmt.Println("failed to setup app: %w", err)
	}
}
