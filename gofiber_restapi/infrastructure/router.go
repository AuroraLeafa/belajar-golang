package app

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/todos")     //Get All
	api.Get("/todos/:id") //Get by Id
	api.Post("/todos")    //Create
	api.Put("/todos/:id") //Update
	api.Delete("/todos/:id")
}
