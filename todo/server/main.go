package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type Todo struct {
	ID        int    `json:"id"`
	Body      string `json:"body"`
	Completed bool   `json:"completed"`
}

func main() {
	fmt.Println("TODO Server Started.")

	app := fiber.New()

	todos := []Todo{}

	app.Get("/healthcheck", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	app.Post("/api/todos", func(c *fiber.Ctx) error {
		todo := &Todo{}

		if err := c.BodyParser(todo); err != nil {
			return err
		}

		todo.ID = len(todos) + 1

		todos = append(todos, *todo)

		return c.JSON(todos)
	})

	app.Patch("/api/todos/:id/completed", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")

		if err != nil {
			return c.Status(401).SendString("Invalid ID")
		}

		for i, t := range todos {
			if t.ID == id {
				todos[i].Completed = !todos[i].Completed
				break
			}
		}

		return c.JSON(todos)
	})

	app.Get("/api/todos", func(c *fiber.Ctx) error {
		return c.JSON(todos)
	})

	log.Fatal(app.Listen(":4000"))
}
