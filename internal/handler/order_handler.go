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

func (h *OrderHandler) Cancel(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return err
	}

	userID := utils.GetUserID(c)

	err = h.Service.Cancel(c.UserContext(), id, userID)
	if err != nil {
		return err
	}

	return utils.Success(
		c,
		fiber.StatusOK,
		"order cancelled",
		nil,
	)
}

func (h *OrderHandler) MarkAsShipped(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return err
	}

	err = h.Service.MarkAsShipped(c.UserContext(), id)
	if err != nil {
		return err
	}

	return utils.Success(
		c,
		fiber.StatusOK,
		"order shipped",
		nil,
	)

}

func (h *OrderHandler) MarkAsCompleted(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return err
	}

	err = h.Service.MarkAsCompleted(c.UserContext(), id)
	if err != nil {
		return err
	}

	return utils.Success(
		c,
		fiber.StatusOK,
		"order completed",
		nil,
	)
}

func (h *OrderHandler) MarkAsPaid(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return err
	}
	err = h.Service.MarkAsPaid(c.UserContext(), id)
	if err != nil {
		return err
	}

	return utils.Success(
		c,
		fiber.StatusOK,
		"payment success",
		nil,
	)
}

func (h *OrderHandler) MarkAsFailed(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return err
	}
	err = h.Service.MarkAsFailed(c.UserContext(), id)
	if err != nil {
		return err
	}

	return utils.Success(
		c,
		fiber.StatusOK,
		"payment failed",
		nil,
	)
}
