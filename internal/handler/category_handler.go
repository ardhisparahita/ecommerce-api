package handler

import (
	"github.com/ardhisparahita/ecommerce-api/internal/dto/request"
	_ "github.com/ardhisparahita/ecommerce-api/internal/dto/response"
	"github.com/ardhisparahita/ecommerce-api/internal/service"
	"github.com/ardhisparahita/ecommerce-api/pkg/utils"
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

// Create godoc
//
// @Summary Create category
// @Description Create a new category
// @Tags Categories
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param request body request.CreateCategoryRequest true "Category Request"
// @Success 201 {object} response.CategorySwaggerResponse
// @Failure 400 {object} response.ErrorSwaggerResponse
// @Failure 401 {object} response.ErrorSwaggerResponse
// @Failure 422 {object} response.ErrorSwaggerResponse
// @Failure 500 {object} response.ErrorSwaggerResponse
// @Router /categories [post]
func (h *CategoryHandler) Create(c *fiber.Ctx) error {
	var req request.CreateCategoryRequest

	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := utils.ValidationStruct(req); err != nil {
		return utils.ResponseError(c, err)
	}

	data, err := h.Service.Create(c.UserContext(), req)
	if err != nil {
		return err
	}

	return utils.ResponseSuccess(
		c,
		fiber.StatusOK,
		"category created",
		data,
	)

}

// FindAll godoc
//
// @Summary Get all categories
// @Description Get all categories
// @Tags Categories
// @Security BearerAuth
// @Produce json
// @Success 200 {object} response.CategoryListSwaggerResponse
// @Failure 401 {object} response.ErrorSwaggerResponse
// @Failure 500 {object} response.ErrorSwaggerResponse
// @Router /categories [get]
func (h *CategoryHandler) FindAll(c *fiber.Ctx) error {
	data, err := h.Service.FindAll(c.UserContext())

	if err != nil {
		return err
	}

	return utils.ResponseSuccess(
		c,
		fiber.StatusOK,
		"get all categories",
		data,
	)
}
