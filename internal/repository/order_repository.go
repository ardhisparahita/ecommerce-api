package repository

import (
	"context"

	"github.com/ardhisparahita/ecommerce-api/internal/domain"
	"gorm.io/gorm"
)

type OrderRepository interface {
	CreateTx(ctx context.Context, tx *gorm.DB, order *domain.Order) error
	FindAllByUserID(ctx context.Context, userID uint64) ([]domain.Order, error)
	FindByIDAndUserID(ctx context.Context, id uint64, userID uint64) (*domain.Order, error)
}
