package service

import (
	"context"

	"github.com/ardhisparahita/ecommerce-api/internal/dto/response"
	"github.com/ardhisparahita/ecommerce-api/internal/repository"
)

type OrderServiceImpl struct {
	Repo repository.OrderRepository
}

func NewOrderService(repo repository.OrderRepository) OrderService {
	return &OrderServiceImpl{
		Repo: repo,
	}
}

func (s *OrderServiceImpl) FindAll(ctx context.Context, userID uint64) ([]response.OrderListResponse, error) {
	orders, err := s.Repo.FindAllByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	var result []response.OrderListResponse

	for _, order := range orders {
		result = append(result, response.OrderListResponse{
			ID:          order.ID,
			TotalAmount: order.TotalAmount,
			Status:      order.Status,
			CreatedAt:   order.CreatedAt,
		})
	}

	return result, nil
}

func (s *OrderServiceImpl) FindByID(ctx context.Context, id uint64, userID uint64) (*response.OrderDetailResponse, error) {
	order, err := s.Repo.FindByIDAndUserID(ctx, id, userID)
	if err != nil {
		return nil, err
	}

	var items []response.OrderItemResponse

	for _, item := range order.OrderItems {
		items = append(items, response.OrderItemResponse{
			ProductID:   item.ProductID,
			ProductName: item.ProductName,
			Price:       item.Price,
			Quantity:    item.Quantity,
			Subtotal:    item.Subtotal,
		})
	}

	return &response.OrderDetailResponse{
		ID:            order.ID,
		RecipientName: order.RecipientName,
		Phone:         order.Phone,
		Address:       order.Address,
		City:          order.City,
		PostalCode:    order.PostalCode,
		TotalAmount:   order.TotalAmount,
		Status:        order.Status,
		Items:         items,
		Payment: response.PaymentDetailResponse{
			Method: order.Payment.Method,
			Status: order.Status,
			Amount: order.Payment.Amount,
		},
	}, nil
}
