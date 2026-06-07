package repository

import (
	"context"

	"github.com/ardhisparahita/ecommerce-api/internal/domain"
	"github.com/ardhisparahita/ecommerce-api/internal/dto/request"
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

func (r *ProductRepositoryImpl) FindAll(ctx context.Context, req request.ProductQueryRequest) ([]domain.Product, int64, error) {
	var products []domain.Product
	var totalRows int64

	db := r.DB.WithContext(ctx).Model(&domain.Product{}).Preload("Category")

	if req.Search != "" {
		db = db.Where("name LIKE ?", "%"+req.Search+"%")
	}

	if req.CategoryID > 0 {
		db = db.Where("category_id = ?", req.CategoryID)
	}

	if err := db.Count(&totalRows).Error; err != nil {
		return nil, 0, err
	}

	offset := (req.Page - 1) * req.Limit

	err := db.Offset(offset).Limit(req.Limit).Find(&products).Error
	if err != nil {
		return nil, 0, err
	}

	return products, totalRows, err
}

func (r *ProductRepositoryImpl) FindByID(ctx context.Context, id uint64) (*domain.Product, error) {
	var product domain.Product

	err := r.DB.WithContext(ctx).Preload("Category").First(&product, id).Error

	return &product, err
}

func (r *ProductRepositoryImpl) Update(ctx context.Context, product *domain.Product) error {
	return r.DB.WithContext(ctx).Save(product).Error
}

func (r *ProductRepositoryImpl) UpdateTx(ctx context.Context, tx *gorm.DB, product *domain.Product) error {
	return tx.WithContext(ctx).Save(product).Error
}

func (r *ProductRepositoryImpl) Delete(ctx context.Context, id uint64) error {
	return r.DB.WithContext(ctx).Delete(&domain.Product{}, id).Error
}
