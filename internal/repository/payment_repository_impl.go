package repository

import (
	"context"

	"github.com/ardhisparahita/ecommerce-api/internal/domain"
	"gorm.io/gorm"
)

type PaymentRepositoryImpl struct {
	DB *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) PaymentRepository {
	return &PaymentRepositoryImpl{
		DB: db,
	}
}

func (r *PaymentRepositoryImpl) CreateTx(ctx context.Context, tx *gorm.DB, payment *domain.Payment) error {
	return tx.WithContext(ctx).Create(payment).Error
}
