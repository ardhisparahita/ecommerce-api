package service

import (
	"context"

	"github.com/ardhisparahita/ecommerce-api/internal/dto/response"
)

type OrderService interface {
	FindAll(ctx context.Context, userID uint64) ([]response.OrderListResponse, error)
	FindByID(ctx context.Context, id uint64, userID uint64) (*response.OrderDetailResponse, error)
}
