package main

import (
	"test_project/handler/it09handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupITRoutes(api fiber.Router) {

	it09Handler := it09handler.NewCommentHandler()
	it09Group := api.Group("it09")
	it09Group.
		Get("/comments", it09Handler.GetComments).
		Post("/insert-comment", it09Handler.InsertComment)

}

func main() {

	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New())

	api := app.Group("/api")
	SetupITRoutes(api)

	app.Listen(":3000")
}
