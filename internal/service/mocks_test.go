package service

import (
	"context"

	"github.com/ardhisparahita/ecommerce-api/internal/domain"
	"github.com/ardhisparahita/ecommerce-api/internal/dto/request"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type MockOrderRepository struct {
	mock.Mock
}

func (m *MockOrderRepository) CreateTx(
	ctx context.Context,
	tx *gorm.DB,
	order *domain.Order,
) error {

	args := m.Called(ctx, tx, order)
	return args.Error(0)
}

func (m *MockOrderRepository) UpdateTx(
	ctx context.Context,
	tx *gorm.DB,
	order *domain.Order,
) error {

	args := m.Called(ctx, tx, order)
	return args.Error(0)
}

func (m *MockOrderRepository) FindByID(
	ctx context.Context,
	id uint64,
) (*domain.Order, error) {

	args := m.Called(ctx, id)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*domain.Order), args.Error(1)
}

func (m *MockOrderRepository) Update(
	ctx context.Context,
	order *domain.Order,
) error {

	args := m.Called(ctx, order)
	return args.Error(0)
}

func (m *MockOrderRepository) FindAllByUserID(
	ctx context.Context,
	userID uint64,
) ([]domain.Order, error) {

	args := m.Called(ctx, userID)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).([]domain.Order), args.Error(1)
}

func (m *MockOrderRepository) FindByIDAndUserID(
	ctx context.Context,
	id uint64,
	userID uint64,
) (*domain.Order, error) {

	args := m.Called(ctx, id, userID)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*domain.Order), args.Error(1)
}

func (m *MockOrderRepository) FindByIDWithItems(
	ctx context.Context,
	id uint64,
) (*domain.Order, error) {

	args := m.Called(ctx, id)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*domain.Order), args.Error(1)
}

func (m *MockOrderRepository) FindByIDAndUserIDWithItems(
	ctx context.Context,
	id uint64,
	userID uint64,
) (*domain.Order, error) {

	args := m.Called(ctx, id, userID)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*domain.Order), args.Error(1)
}

type MockProductRepository struct {
	mock.Mock
}

func (m *MockProductRepository) Create(
	ctx context.Context,
	product *domain.Product,
) error {

	args := m.Called(ctx, product)
	return args.Error(0)
}

func (m *MockProductRepository) FindAll(
	ctx context.Context,
	req request.ProductQueryRequest,
) ([]domain.Product, int64, error) {

	args := m.Called(ctx, req)

	var products []domain.Product
	if args.Get(0) != nil {
		products = args.Get(0).([]domain.Product)
	}

	return products, args.Get(1).(int64), args.Error(2)
}

func (m *MockProductRepository) FindByID(
	ctx context.Context,
	id uint64,
) (*domain.Product, error) {

	args := m.Called(ctx, id)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*domain.Product), args.Error(1)
}

func (m *MockProductRepository) Update(
	ctx context.Context,
	product *domain.Product,
) error {

	args := m.Called(ctx, product)
	return args.Error(0)
}

func (m *MockProductRepository) UpdateTx(
	ctx context.Context,
	tx *gorm.DB,
	product *domain.Product,
) error {

	args := m.Called(ctx, tx, product)
	return args.Error(0)
}

func (m *MockProductRepository) Delete(
	ctx context.Context,
	id uint64,
) error {

	args := m.Called(ctx, id)
	return args.Error(0)
}

type MockPaymentRepository struct {
	mock.Mock
}

func (m *MockPaymentRepository) CreateTx(
	ctx context.Context,
	tx *gorm.DB,
	payment *domain.Payment,
) error {

	args := m.Called(ctx, tx, payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FIndByOrderID(
	ctx context.Context,
	orderID uint64,
) (*domain.Payment, error) {

	args := m.Called(ctx, orderID)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*domain.Payment), args.Error(1)
}

func (m *MockPaymentRepository) UpdateTx(
	ctx context.Context,
	tx *gorm.DB,
	payment *domain.Payment,
) error {

	args := m.Called(ctx, tx, payment)
	return args.Error(0)
}

type MockCategoryRepository struct {
	mock.Mock
}

func (m *MockCategoryRepository) Create(ctx context.Context, category *domain.Category) error {
	args := m.Called(ctx, category)

	return args.Error(0)
}

func (m *MockCategoryRepository) FindAll(ctx context.Context) ([]domain.Category, error) {
	args := m.Called(ctx)

	return args.Get(0).([]domain.Category), args.Error(1)
}

func (m *MockCategoryRepository) FindByID(ctx context.Context, id uint64) (*domain.Category, error) {
	args := m.Called(ctx, id)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*domain.Category), args.Error(1)
}

func (m *MockCategoryRepository) Update(ctx context.Context, category *domain.Category) error {
	args := m.Called(ctx, category)

	return args.Error(0)
}
