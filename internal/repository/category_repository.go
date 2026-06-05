package repository

import (
	"context"

	"github.com/ardhisparahita/ecommerce-api/internal/domain"
)

type CategoryRepository interface {
	Create(ctx context.Context, category *domain.Category) error
	FindAll(ctx context.Context) ([]domain.Category, error)
}
