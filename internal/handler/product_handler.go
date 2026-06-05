package handler

import (
	"strconv"

	"github.com/ardhisparahita/ecommerce-api/internal/dto/request"
	"github.com/ardhisparahita/ecommerce-api/internal/service"
	"github.com/gofiber/fiber/v2"
)

type ProductHandler struct {
	Service service.ProductService
}

func NewProductHandler(service service.ProductService) *ProductHandler {
	return &ProductHandler{
		Service: service,
	}
}

func (h *ProductHandler) Create(c *fiber.Ctx) error {
	var req request.CreateProductRequest

	if err := c.BodyParser(&req); err != nil {
		return err
	}

	err := h.Service.Create(c.UserContext(), req)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"message": "product created",
	})
}

func (h *ProductHandler) FindAll(c *fiber.Ctx) error {
	res, err := h.Service.FindAll(c.UserContext())
	if err != nil {
		return err
	}

	return c.JSON(res)
}

func (h *ProductHandler) FindByID(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadGateway,
			"invalid id",
		)
	}

	res, err := h.Service.FindByID(c.UserContext(), id)
	if err != nil {
		return err
	}

	return c.JSON(res)
}

func (h *ProductHandler) Update(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadGateway,
			"invalid id",
		)
	}

	var req request.UpdateProductRequest

	if err := c.BodyParser(&req); err != nil {
		return err
	}

	err = h.Service.Update(c.UserContext(), id, req)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"message": "product updated",
	})
}

func (h *ProductHandler) Delete(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			"invalid id",
		)
	}

	err = h.Service.Delete(c.UserContext(), id)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"message": "product deleted",
	})
}
