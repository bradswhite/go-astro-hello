package main

import (
	"log"
	"os"
	_ "github.com/joho/godotenv/autoload"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return ":" + port
}


func main() {
  app := fiber.New()

  app.Use(cors.New(cors.Config{
    AllowHeaders:     "Content-Type,access-control-allow-origin, access-control-allow-headers",
    AllowOrigins:     "*",
  }))

	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Astro + GO")
	})

	log.Fatal(app.Listen(getPort()))
}
