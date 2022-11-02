package routes

import (
	"github.com/gofiber/fiber/v2"
	"go_crud/db"
	"go_crud/db/models"
	"go_crud/services"
)

func SetupPersonRoutes(app *fiber.App, database db.Database) {
	app.Get("/person", func(fiberContext *fiber.Ctx) error {
		personList, err := services.GetPersonService(database)
		if err != nil {
			return fiberContext.Status(500).JSON(fiber.Map{"message": "Error getting persons"})
		}
		return fiberContext.Status(200).JSON(personList)
	})

	app.Post("/person", func(fiberContext *fiber.Ctx) error {
		var person = new(models.Person)
		err := fiberContext.BodyParser(person)
		if err != nil {
			return fiberContext.Status(400).JSON(fiber.Map{"message": "Error parsing person"})
		}
		createPersonError := services.CreatePersonService(person, database)
		if createPersonError != nil {
			return fiberContext.Status(400).JSON(fiber.Map{"message": "Error creating person", "error": createPersonError.Error()})
		}
		return fiberContext.Status(200).JSON(person)
	})
}
