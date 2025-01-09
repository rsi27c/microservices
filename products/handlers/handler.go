package handlers

import (
	"example/productservices/pkg/models"
	"example/productservices/services"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	services.Services
}

func (handler *Handler) CreateProduct(ctx *fiber.Ctx) error {
	var product models.Products

	if err := ctx.BodyParser(&product); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err)
	}

	if err := handler.Services.CreateProduct(&product); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(err)
	}

	return ctx.Status(http.StatusCreated).JSON(&product)
}

func (handler *Handler) GetProduct(ctx *fiber.Ctx) error {
	product, err := handler.Services.GetProduct()
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(err)
	}

	return ctx.Status(http.StatusOK).JSON(&product)
}

func (handler *Handler) GetProductByID(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err)
	}

	product, err := handler.Services.GetProductByID(id)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(err)
	}

	return ctx.Status(http.StatusOK).JSON(&product)
}

func (handler *Handler) UpdateProduct(ctx *fiber.Ctx) error {
	var product models.Products

	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err)
	}

	if err := ctx.BodyParser(&product); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err)
	}

	if err := handler.Services.UpdateProduct(&product, id); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(err)
	}

	return ctx.Status(http.StatusCreated).JSON(&product)
}

func (handler *Handler) DeleteProduct(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err)
	}

	err = handler.Services.DeleteProduct(id)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(err)
	}

	return ctx.Status(http.StatusOK).JSON(id)
}
