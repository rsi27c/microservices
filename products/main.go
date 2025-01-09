package main

import (
	"example/productservices/routes"
	"example/productservices/internals"
	"example/productservices/pkg/loggers"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func main() {
	internals.GetEnv()
	loggers.OpenLog()
	db := internals.Connect()
	internals.Migrate(db)

	server := fiber.New()

	routes.GetRoutes(server, db)
	if err := server.Listen(os.Getenv("HTTP_PORT")); err != nil {
		log.Fatalf("Failed to listen to the server %v", err)
	}
}
