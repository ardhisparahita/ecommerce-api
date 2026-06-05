package handler

import (
	"github.com/ardhisparahita/ecommerce-api/internal/dto/request"
	"github.com/ardhisparahita/ecommerce-api/internal/service"
	"github.com/gofiber/fiber/v2"
)

type CategoryHandler struct {
	Service service.CategoryService
}

func NewCategoryHandler(service service.CategoryService) *CategoryHandler {
	return &CategoryHandler{
		Service: service,
	}
}

func (h *CategoryHandler) Create(c *fiber.Ctx) error {
	var req request.CreateCategoryRequest

	if err := c.BodyParser(&req); err != nil {
		return err
	}

	err := h.Service.Create(c.Context(), req)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"message": "Category Created",
	})
}

func (h *CategoryHandler) FindAll(c *fiber.Ctx) error {
	res, err := h.Service.FindAll(c.Context())

	if err != nil {
		return err
	}

	return c.JSON(res)
}
