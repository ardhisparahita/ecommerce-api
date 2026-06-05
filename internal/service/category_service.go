package service

import (
	"context"

	"github.com/ardhisparahita/ecommerce-api/internal/dto/request"
	"github.com/ardhisparahita/ecommerce-api/internal/dto/response"
)

type CategoryService interface {
	Create(ctx context.Context, req request.CreateCategoryRequest) (*response.CategoryResponse, error)
	FindAll(ctx context.Context) ([]response.CategoryResponse, error)
}
