package main

import (
	"log"
	"leadreverse/middleware"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/keyauth"
)


func main() {
    app := fiber.New()

    app.Use(keyauth.New(keyauth.Config{
        KeyLookup:  "header:X-API-Key",
        Validator:  middleware.ValidateAPIKey,
    }))

    // Define a route for the GET method on the root path '/'
    app.Get("/", func(c fiber.Ctx) error {
        // Send a string response to the client
        return c.SendString("Hello, World 👋!")
    })

		log.Println("Server is running on port 3000")

    // Start the server on port 3000
    log.Fatal(app.Listen(":3000"))
}
