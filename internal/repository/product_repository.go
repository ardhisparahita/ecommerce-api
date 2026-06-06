package repository

import (
	"context"

	"github.com/ardhisparahita/ecommerce-api/internal/domain"
	"gorm.io/gorm"
)

type ProductRepository interface {
	Create(ctx context.Context, product *domain.Product) error
	FindAll(ctx context.Context) ([]domain.Product, error)
	FindByID(ctx context.Context, id uint64) (*domain.Product, error)
	Update(ctx context.Context, product *domain.Product) error
	UpdateTx(ctx context.Context, tx *gorm.DB, product *domain.Product) error
	Delete(ctx context.Context, id uint64) error
}
