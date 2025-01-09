package routes

import (
	"example/productservices/handlers"
	"example/productservices/repository"
	"example/productservices/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetRoutes(server *fiber.App, db *gorm.DB) {
	repo := repository.InitRepository(db)
	service := services.InitServices(repo)
	handler := handlers.Handler{Services: service}

	server.Post("/create", handler.CreateProduct)
	server.Get("/get", handler.GetProduct)
	server.Get("/get/:id", handler.GetProductByID)
	server.Put("/update/:id", handler.UpdateProduct)
	server.Delete("/delete/:id", handler.DeleteProduct)
}
