package main

import (
	"github.com/gofiber/fiber/v2"
	"go_crud/db"
	"go_crud/routes"
	"log"
	"os"
)

func main() {
	app := fiber.New()
	databaseConnection, _ := db.InitDB()
	routes.SetupPersonRoutes(app, databaseConnection)
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Fatalln(app.Listen(":" + port))
}
