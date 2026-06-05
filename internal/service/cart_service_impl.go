package service

import (
	"context"
	"errors"

	"github.com/ardhisparahita/ecommerce-api/internal/domain"
	"github.com/ardhisparahita/ecommerce-api/internal/dto/request"
	"github.com/ardhisparahita/ecommerce-api/internal/dto/response"
	"github.com/ardhisparahita/ecommerce-api/internal/mapper"
	"github.com/ardhisparahita/ecommerce-api/internal/repository"
	"gorm.io/gorm"
)

type CartServiceImpl struct {
	Repo repository.CartRepository
}

func NewCartService(repo repository.CartRepository) CartService {
	return &CartServiceImpl{
		Repo: repo,
	}
}

func (s *CartServiceImpl) AddToCart(ctx context.Context, userID uint64, req request.AddToCartRequest) (*response.CartResponse, error) {
	cart, err := s.Repo.FindByUserIDAndProductID(ctx, userID, req.ProductID)

	if err == nil {
		cart.Quantity += req.Quantity

		if err := s.Repo.Update(ctx, cart); err != nil {
			return nil, err
		}

		cart, err := s.Repo.FindByIDAndUserID(ctx, cart.ID, userID)
		if err != nil {
			return nil, err
		}

		return mapper.ToCartResponse(cart), nil
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	newCart := domain.Cart{
		UserID:    userID,
		ProductID: req.ProductID,
		Quantity:  req.Quantity,
	}

	if err := s.Repo.Create(ctx, &newCart); err != nil {
		return nil, err
	}

	cart, err = s.Repo.FindByIDAndUserID(ctx, newCart.ID, userID)
	if err != nil {
		return nil, err
	}

	return mapper.ToCartResponse(cart), nil
}

func (s *CartServiceImpl) FindAll(ctx context.Context, userID uint64) (*response.CartListResponse, error) {
	carts, err := s.Repo.FindAllByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	items, grandTotal := mapper.ToCartResponses(carts)

	return &response.CartListResponse{
		Items:      items,
		TotalItems: len(items),
		GrandTotal: grandTotal,
	}, nil
}

func (s *CartServiceImpl) Update(ctx context.Context, id uint64, userID uint64, req request.UpdateCartRequest) (*response.CartResponse, error) {
	cart, err := s.Repo.FindByIDAndUserID(ctx, id, userID)
	if err != nil {
		return nil, err
	}

	cart.Quantity = req.Quantity

	if err := s.Repo.Update(ctx, cart); err != nil {
		return nil, err
	}

	cart, err = s.Repo.FindByIDAndUserID(ctx, id, userID)
	if err != nil {
		return nil, err
	}

	return mapper.ToCartResponse(cart), nil

}

func (s *CartServiceImpl) Delete(ctx context.Context, id uint64, userID uint64) error {
	_, err := s.Repo.FindByIDAndUserID(ctx, id, userID)
	if err != nil {
		return err
	}
	
	return s.Repo.Delete(ctx, id, userID)
}
