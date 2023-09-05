package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

type Comment struct {
	Title string
	Body  string
}

type Storage struct {
	Comments []Comment
}

var comments = []Comment{
	{
		"Comment #1",
		"Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
	},
	{
		"Comment #2",
		"Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
	},
	{
		"Comment #3",
		"Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
	},
}

var storage = Storage{comments}

func main() {

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", storage)
	})

	log.Fatal(app.Listen(":3000"))
}
