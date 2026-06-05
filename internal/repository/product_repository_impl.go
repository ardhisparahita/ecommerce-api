package repository

import (
	"context"

	"github.com/ardhisparahita/ecommerce-api/internal/domain"
	"gorm.io/gorm"
)

type ProductRepositoryImpl struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &ProductRepositoryImpl{
		DB: db,
	}
}

func (r *ProductRepositoryImpl) Create(ctx context.Context, product *domain.Product) error {
	return r.DB.WithContext(ctx).Create(product).Error
}

func (r *ProductRepositoryImpl) FindAll(ctx context.Context) ([]domain.Product, error) {
	var products []domain.Product

	err := r.DB.WithContext(ctx).Preload("Category").Find(&products).Error
	return products, err
}

func (r *ProductRepositoryImpl) FindByID(ctx context.Context, id uint64) (*domain.Product, error) {
	var product domain.Product

	err := r.DB.WithContext(ctx).Preload("Category").First(&product, id).Error

	return &product, err
}

func (r *ProductRepositoryImpl) Update(ctx context.Context, product *domain.Product) error {
	return r.DB.WithContext(ctx).Updates(product).Error
}

func (r *ProductRepositoryImpl) Delete(ctx context.Context, id uint64) error {
	return r.DB.WithContext(ctx).Delete(&domain.Product{}, id).Error
}
