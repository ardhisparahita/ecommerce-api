package repository

import (
	"context"

	"github.com/ardhisparahita/ecommerce-api/internal/domain"
	"gorm.io/gorm"
)

type CategoryRepositoryImpl struct {
	DB *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &CategoryRepositoryImpl{
		DB: db,
	}
}

func (r *CategoryRepositoryImpl) Create(ctx context.Context, category *domain.Category) error {
	return r.DB.WithContext(ctx).Create(category).Error
}

func (r *CategoryRepositoryImpl) FindAll(ctx context.Context) ([]domain.Category, error) {
	var categories []domain.Category

	err := r.DB.WithContext(ctx).Find(&categories).Error

	return categories, err
}
