package handler

import (
	"github.com/ardhisparahita/ecommerce-api/internal/dto/request"
	"github.com/ardhisparahita/ecommerce-api/internal/service"
	"github.com/ardhisparahita/ecommerce-api/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

type CheckoutHandler struct {
	Service service.CheckoutService
}

func NewCheckoutHandler(service service.CheckoutService) *CheckoutHandler {
	return &CheckoutHandler{
		Service: service,
	}
}

func (h *CheckoutHandler) Checkout(c *fiber.Ctx) error {
	var req request.CheckoutRequest

	if err := c.BodyParser(&req); err != nil {
		return err
	}

	userID := utils.GetUserID(c)

	data, err := h.Service.Checkout(c.UserContext(), userID, req)
	if err != nil {
		return err
	}

	return utils.Success(
		c,
		fiber.StatusOK,
		"checkout success",
		data,
	)
}
