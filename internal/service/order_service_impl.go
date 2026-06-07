package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/ardhisparahita/ecommerce-api/internal/dto/response"
	"github.com/ardhisparahita/ecommerce-api/internal/repository"
	"gorm.io/gorm"
)

type OrderServiceImpl struct {
	DB          *gorm.DB
	OrderRepo   repository.OrderRepository
	ProductRepo repository.ProductRepository
	PaymentRepo repository.PaymentRepository
}

func NewOrderService(
	db *gorm.DB,
	orderRepo repository.OrderRepository,
	productRepo repository.ProductRepository,
	paymentRepo repository.PaymentRepository,
) OrderService {
	return &OrderServiceImpl{
		DB:          db,
		OrderRepo:   orderRepo,
		ProductRepo: productRepo,
		PaymentRepo: paymentRepo,
	}
}

func (s *OrderServiceImpl) FindAll(ctx context.Context, userID uint64) ([]response.OrderListResponse, error) {
	orders, err := s.OrderRepo.FindAllByUserID(ctx, userID)
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
	order, err := s.OrderRepo.FindByIDAndUserID(ctx, id, userID)
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

func (s *OrderServiceImpl) MarkAsPaid(ctx context.Context, id uint64) error {
	order, err := s.OrderRepo.FindByID(ctx, id)
	if err != nil {
		return err
	}

	fmt.Println("ORDER ID:", order.ID)
	fmt.Println("ORDER STATUS:", order.Status)

	if order.Status != "PENDING" {
		return errors.New("order already processed")
	}

	payment, err := s.PaymentRepo.FIndByOrderID(ctx, order.ID)
	if err != nil {
		return err
	}

	return s.DB.Transaction(func(tx *gorm.DB) error {
		payment.Status = "PAID"

		if err := s.PaymentRepo.UpdateTx(ctx, tx, payment); err != nil {
			return err
		}

		order.Status = "PAID"

		if err := s.OrderRepo.UpdateTx(ctx, tx, order); err != nil {
			return err
		}

		return nil
	})
}

func (s *OrderServiceImpl) MarkAsFailed(ctx context.Context, id uint64) error {
	order, err := s.OrderRepo.FindByID(ctx, id)
	if err != nil {
		return err
	}

	if order.Status != "PENDING" {
		return errors.New("order already processed")
	}

	payment, err := s.PaymentRepo.FIndByOrderID(ctx, order.ID)
	if err != nil {
		return err
	}

	return s.DB.Transaction(func(tx *gorm.DB) error {
		for _, item := range order.OrderItems {
			product, err := s.ProductRepo.FindByID(ctx, item.ProductID)
			if err != nil {
				return err
			}

			product.Stock += item.Quantity

			if err := s.ProductRepo.UpdateTx(ctx, tx, product); err != nil {
				return err
			}
		}

		payment.Status = "FAILED"

		if err := s.OrderRepo.UpdateTx(ctx, tx, order); err != nil {
			return err
		}

		return nil
	})
}

func (s *OrderServiceImpl) Cancel(ctx context.Context, id uint64, userID uint64) error {
	order, err := s.OrderRepo.FindByIDAndUserID(ctx, id, userID)
	if err != nil {
		return err
	}

	if order.Status != "PENDING" {
		return errors.New(
			"only pending orders can be cancelled",
		)
	}

	payment, err := s.PaymentRepo.FIndByOrderID(ctx, order.ID)
	if err != nil {
		return err
	}

	return s.DB.Transaction(func(tx *gorm.DB) error {
		for _, item := range order.OrderItems {
			product, err := s.ProductRepo.FindByID(ctx, item.ProductID)
			if err != nil {
				return err
			}

			product.Stock += item.Quantity

			if err := s.ProductRepo.UpdateTx(ctx, tx, product); err != nil {
				return err
			}
		}

		order.Status = "CANCELLED"

		if err := s.OrderRepo.UpdateTx(ctx, tx, order); err != nil {
			return err
		}

		payment.Status = "FAILED"

		if err := s.PaymentRepo.UpdateTx(ctx, tx, payment); err != nil {
			return err
		}

		return nil
	})
}

func (s *OrderServiceImpl) MarkAsShipped(ctx context.Context, id uint64) error {
	order, err := s.OrderRepo.FindByID(ctx, id)
	if err != nil {
		return err
	}

	if order.Status != "PAID" {
		return errors.New("Order must be paid before shipping")
	}

	order.Status = "SHIPPED"

	return s.OrderRepo.Update(ctx, order)
}

func (s *OrderServiceImpl) MarkAsCompleted(ctx context.Context, id uint64) error {
	order, err := s.OrderRepo.FindByID(ctx, id)
	if err != nil {
		return err
	}

	if order.Status != "SHIPPED" {
		return errors.New("Order must be shipped first")
	}

	order.Status = "COMPLETED"

	return s.OrderRepo.Update(ctx, order)
}
