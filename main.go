package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	app.Use(cors.New())

	app.Static("/", "./client/dist")

	app.Get("/healthchecker", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"status":  "success",
			"message": "Welcome to Golang, Fiber and Mongo DB",
		})
	})

	app.Get("/users", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"data": "User List from the backend",
		})
	})

	log.Println("🚀 Starting up on port 5500")
	app.Listen(":5500")
}

/*
https://github.com/FaztWeb/go-react-project

https://github.com/gofiber/fiber/blob/master/.github/README_es.md
https://docs.gofiber.io/
https://gofiber.io/

https://www.mongodb.com/docs/drivers/go/current/fundamentals/crud/write-operations/delete/
https://github.com/mongodb/mongo-go-driver

https://codevoweb.com/build-a-simple-api-in-golang-using-fiber-and-sqlite/
https://github.com/wpcodevo/golang-fiber-sqlite

https://github.com/Prakashdeveloper03/Fiber-todo-app
*/
