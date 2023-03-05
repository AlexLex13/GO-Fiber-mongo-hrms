package main

import "github.com/gofiber/fiber"

type MongoInstance struct {
	Client
	Db
}

var mg MongoInstance

const dbName = "fiber-hrms"
const mongoURI = "mongodb://localhost:27017" + dbName

func main() {
	app := fiber.New()

	app.Get("/employee", func(c *fiber.Ctx) {

	})
	app.Post("/employee")
	app.Put("/employee/:id")
	app.Delete("/employee/:id")
}
