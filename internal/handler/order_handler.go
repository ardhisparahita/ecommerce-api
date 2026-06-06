package handler

import (
	"strconv"

	"github.com/ardhisparahita/ecommerce-api/internal/service"
	"github.com/ardhisparahita/ecommerce-api/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

type OrderHandler struct {
	Service service.OrderService
}

func NewOrderHandler(service service.OrderService) *OrderHandler {
	return &OrderHandler{Service: service}
}

func (h *OrderHandler) FindAll(c *fiber.Ctx) error {
	userID := utils.GetUserID(c)

	data, err := h.Service.FindAll(c.UserContext(), userID)
	if err != nil {
		return err
	}

	return utils.Success(
		c,
		fiber.StatusOK,
		"get orders",
		data,
	)
}

func (h *OrderHandler) FindByID(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return err
	}

	userID := utils.GetUserID(c)

	data, err := h.Service.FindByID(c.UserContext(), id, userID)
	if err != nil {
		return err
	}

	return utils.Success(
		c,
		fiber.StatusOK,
		"get order details",
		data,
	)
}
