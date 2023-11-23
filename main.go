package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"time"
)

type Todos struct {
	ID    int    
	Title string `json:"title"`
	Todo  string `json:"todo"`
	Do    bool   `json:"do"`
}

func randomID() int {
	crtime := time.Now()
	id := crtime.Minute() + crtime.Second() + crtime.Hour() + crtime.Day() + crtime.Year() + int(crtime.Month())
	return id
}

func addTodo(c *fiber.Ctx) error {
	var td Todos
	if err := c.BodyParser(&td); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request body")
	}

	for _, t := range list {
		if td.Title == t.Title && !t.Do {
			return c.Status(fiber.StatusConflict).SendString("You already have this todo")
		}
	}

	newTodo := Todos{ID: randomID(), Title: td.Title, Todo: td.Todo, Do: td.Do}
	list = append(list, newTodo)

	return c.Status(fiber.StatusCreated).JSON(newTodo)
}

var list []Todos

func main() {
	app := fiber.New()
	fmt.Println("The server is running on port 8080")

	app.Post("/todos", addTodo)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Home page")
	})

	app.Listen(":8080")
}
