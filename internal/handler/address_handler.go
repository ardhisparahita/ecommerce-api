package handler

import (
	"strconv"

	"github.com/ardhisparahita/ecommerce-api/internal/dto/request"
	"github.com/ardhisparahita/ecommerce-api/internal/service"
	"github.com/ardhisparahita/ecommerce-api/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

type AddressHandler struct {
	Service service.AddressService
}

func NewAddressHandler(service service.AddressService) *AddressHandler {
	return &AddressHandler{Service: service}
}

func (h *AddressHandler) Create(c *fiber.Ctx) error {
	var req request.CreateAddressRequest

	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := utils.ValidationStruct(req); err != nil {
		return utils.ResponseError(c, err)
	}

	userID := utils.GetUserID(c)

	data, err := h.Service.Create(c.UserContext(), userID, req)
	if err != nil {
		return err
	}

	return utils.ResponseSuccess(
		c,
		fiber.StatusOK,
		"address created",
		data,
	)
}

func (h *AddressHandler) FindAll(c *fiber.Ctx) error {
	userID := utils.GetUserID(c)

	addresses, err := h.Service.FindAllByUserID(c.UserContext(), userID)
	if err != nil {
		return err
	}

	return utils.ResponseSuccess(
		c,
		fiber.StatusOK,
		"get all addresses",
		addresses,
	)
}

func (h *AddressHandler) FindByID(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid address Id")
	}

	userID := utils.GetUserID(c)

	address, err := h.Service.FindByID(c.UserContext(), id, userID)
	if err != nil {
		return err
	}

	return utils.ResponseSuccess(
		c,
		fiber.StatusOK,
		"get one address",
		address,
	)
}

func (h *AddressHandler) Update(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid address Id")
	}

	var req request.UpdateAddressRequest

	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := utils.ValidationStruct(req); err != nil {
		return utils.ResponseError(c, err)
	}

	userID := utils.GetUserID(c)

	data, err := h.Service.Update(c.UserContext(), id, userID, req)
	if err != nil {
		return err
	}

	return utils.ResponseSuccess(
		c,
		fiber.StatusOK,
		"address updated",
		data,
	)
}

func (h *AddressHandler) Delete(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid address Id")
	}

	userID := utils.GetUserID(c)

	err = h.Service.Delete(c.UserContext(), id, userID)
	if err != nil {
		return err
	}

	return utils.ResponseSuccess(
		c,
		fiber.StatusOK,
		"address deleted",
		nil,
	)
}
