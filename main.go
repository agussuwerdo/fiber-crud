package main

import (
	"fiber-crud/handlers"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Initialize MongoDB and handle errors
	mongoClient, err := handlers.InitMongoDB()
	if err != nil {
		log.Fatalf("Failed to initialize MongoDB: %v", err)
	}

	// Get database name from environment
	dbName := os.Getenv("MONGO_DB")
	if dbName == "" {
		dbName = "fibercrud" // Default database name if not set
	}
	itemsCollection := "items"
	// Pass the MongoDB client to handlers
	handlers.SetClient(mongoClient, dbName, itemsCollection)

	// Routes
	app.Post("/login", handlers.Login)
	app.Get("/items", handlers.GetItems)    // Assuming auth middleware is not needed for now
	app.Post("/items", handlers.CreateItem) // Assuming auth middleware is not needed for now
	app.Put("/items/:id", handlers.UpdateItem)
	app.Delete("/items/:id", handlers.DeleteItem)
	appPort := os.Getenv("PORT")
	if appPort == "" {
		appPort = "3000" // default to port 3000 if PORT is not set
	}

	log.Fatal(app.Listen(":" + appPort))
}
