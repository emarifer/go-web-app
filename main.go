package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	Id   primitive.ObjectID `json:"_id" bson:"_id"`
	Name string             `json:"name"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	mongodbAtlasUri := os.Getenv("MONGODB_ATLAS_URI")
	mongodbName := os.Getenv("MONGO_INITDB_DATABASE")

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongodbAtlasUri))
	if err != nil {
		panic(err)
	}

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
		var users []User

		coll := client.Database(mongodbName).Collection("users")
		// Ordenación descendente según campo (en este caso alfabética)
		// https://www.mongodb.com/docs/drivers/go/current/fundamentals/crud/read-operations/sort/#descending
		opts := options.Find().SetSort(bson.D{{Key: "name", Value: -1}})
		result, err := coll.Find(context.TODO(), bson.M{}, opts)
		if err != nil {
			panic(err)
		}

		for result.Next(context.TODO()) {
			var user User
			result.Decode(&user)
			users = append(users, user)
		}

		return c.Status(200).JSON(fiber.Map{
			"data": users,
		})
	})

	app.Post("/users", func(c *fiber.Ctx) error {
		var user User
		c.BodyParser(&user)

		coll := client.Database(mongodbName).Collection("users")
		result, err := coll.InsertOne(context.TODO(), bson.D{{
			Key:   "name",
			Value: user.Name,
		}})
		if err != nil {
			panic(err)
		}

		return c.Status(201).JSON(fiber.Map{
			"data": result,
		})
	})

	port := os.Getenv("PORT")
	log.Printf("🚀 Starting up on port %s", port)
	app.Listen(fmt.Sprintf(":%s", port))
}

/*
https://github.com/FaztWeb/go-react-project

USO DE PROXY EN EL packege.json PARA EVITAR TENER QUE CAMBIAR EN
EL FRONTEND LA URL DE DESARROLLO POR LA DE DEPLOY. VER:
https://youtu.be/zlrnwGZMBbU?si=BaDbpfALXLr3gS0X&t=1662

https://github.com/gofiber/fiber/blob/master/.github/README_es.md
https://docs.gofiber.io/
https://gofiber.io/

https://www.mongodb.com/docs/drivers/go/current/fundamentals/crud/write-operations/delete/
https://github.com/mongodb/mongo-go-driver

https://codevoweb.com/build-a-simple-api-in-golang-using-fiber-and-sqlite/
https://github.com/wpcodevo/golang-fiber-sqlite

https://github.com/Prakashdeveloper03/Fiber-todo-app
*/
