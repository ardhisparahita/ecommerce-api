package handler

import (
	"github.com/ardhisparahita/ecommerce-api/internal/dto/request"
	"github.com/ardhisparahita/ecommerce-api/internal/service"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	Service service.AuthService
}

func NewAuthHandler(service service.AuthService) *AuthHandler {
	return &AuthHandler{
		Service: service,
	}
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var req request.RegisterRequest

	if err := c.BodyParser(&req); err != nil {
		return err
	}

	err := h.Service.Register(c.Context(), req)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"Message": "Register Success",
	})
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var req request.LoginRequest

	if err := c.BodyParser(&req); err != nil {
		return err
	}

	res, err := h.Service.Login(c.Context(), req)
	if err != nil {
		return err
	}

	return c.JSON(res)
}
