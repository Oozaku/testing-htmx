package main

import (
	"hash/fnv"
	"log"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

type Comment struct {
	Title string
	Body  string
}

type Storage struct {
	mu       sync.Mutex
	Comments map[uint32]Comment
}

var storage = Storage{Comments: make(map[uint32]Comment)}

func addComment(title string, body string) {
	storage.mu.Lock()
	defer storage.mu.Unlock()

	hash := fnv.New32a()
	hash.Write([]byte(title + body))
	id := hash.Sum32()

	storage.Comments[id] = Comment{title, body}
}

func main() {

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		storage.mu.Lock()
		defer storage.mu.Unlock()
		return c.Render("index", storage)
	})

	app.Post("/comment", func(c *fiber.Ctx) error {

		addComment(c.FormValue("title"), c.FormValue("body"))

		return c.Render("components/comments", storage)
	})

	log.Fatal(app.Listen(":3000"))
}
