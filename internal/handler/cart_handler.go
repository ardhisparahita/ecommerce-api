package handler

import (
	"strconv"

	"github.com/ardhisparahita/ecommerce-api/internal/dto/request"
	"github.com/ardhisparahita/ecommerce-api/internal/service"
	"github.com/ardhisparahita/ecommerce-api/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

type CartHandler struct {
	Service service.CartService
}

func NewCartHandler(service service.CartService) *CartHandler {
	return &CartHandler{Service: service}
}

func (h *CartHandler) AddToCart(c *fiber.Ctx) error {
	var req request.AddToCartRequest

	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	userID := utils.GetUserID(c)

	data, err := h.Service.AddToCart(c.UserContext(), userID, req)
	if err != nil {
		return err
	}

	return utils.Success(
		c,
		fiber.StatusOK,
		"product added to cart",
		data,
	)
}

func (h *CartHandler) FindAll(c *fiber.Ctx) error {
	userID := utils.GetUserID(c)

	data, err := h.Service.FindAll(c.UserContext(), userID)
	if err != nil {
		return err
	}

	return utils.Success(
		c,
		fiber.StatusOK,
		"get all carts",
		data,
	)
}

func (h *CartHandler) Update(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid cart id")
	}

	var req request.UpdateCartRequest

	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	userID := utils.GetUserID(c)

	data, err := h.Service.Update(c.UserContext(), id, userID, req)
	if err != nil {
		return err
	}

	return utils.Success(
		c,
		fiber.StatusOK,
		"cart updated",
		data,
	)
}

func (h *CartHandler) Delete(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid cart id")
	}

	userID := utils.GetUserID(c)
	if err := h.Service.Delete(c.UserContext(), id, userID); err != nil {
		return err
	}

	return utils.Success(
		c,
		fiber.StatusOK,
		"cart deleted",
		nil,
	)
}
