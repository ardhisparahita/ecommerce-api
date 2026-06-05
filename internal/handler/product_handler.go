package handler

import (
	"strconv"

	"github.com/ardhisparahita/ecommerce-api/internal/dto/request"
	"github.com/ardhisparahita/ecommerce-api/internal/service"
	"github.com/ardhisparahita/ecommerce-api/pkg/utils"
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
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	data, err := h.Service.Create(c.UserContext(), req)
	if err != nil {
		return err
	}

	return utils.Success(
		c,
		fiber.StatusOK,
		"product created",
		data,
	)
}

func (h *ProductHandler) FindAll(c *fiber.Ctx) error {
	res, err := h.Service.FindAll(c.UserContext())
	if err != nil {
		return err
	}

	return utils.Success(
		c,
		fiber.StatusOK,
		"get all products",
		res,
	)
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

	return utils.Success(
		c,
		fiber.StatusOK,
		"get one product",
		res,
	)
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

	data, err := h.Service.Update(c.UserContext(), id, req)
	if err != nil {
		return err
	}

	return utils.Success(
		c,
		fiber.StatusOK,
		"product updated",
		data,
	)
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

	return utils.Success(
		c,
		fiber.StatusOK,
		"product deleted",
		nil,
	)
}
