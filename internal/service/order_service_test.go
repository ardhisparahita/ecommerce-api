package service

import (
	"context"
	"testing"

	"github.com/ardhisparahita/ecommerce-api/internal/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type MockOrderRepository struct {
	mock.Mock
}

func (m *MockOrderRepository) FindAllByUserID(ctx context.Context, userID uint64) ([]domain.Order, error) {
	args := m.Called(ctx, userID)

	return args.Get(0).([]domain.Order), args.Error(1)
}

func (m *MockOrderRepository) FindByIDAndUserID(ctx context.Context, id uint64, userID uint64) (*domain.Order, error) {
	args := m.Called(ctx, id, userID)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*domain.Order), args.Error(1)
}

func (m *MockOrderRepository) FindByID(ctx context.Context, id uint64) (*domain.Order, error) {
	args := m.Called(ctx, id)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*domain.Order), args.Error(1)
}

func (m *MockOrderRepository) Update(ctx context.Context, order *domain.Order) error {
	args := m.Called(ctx, order)

	return args.Error(0)
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

func TestFindAllOrderSuccess(t *testing.T) {
	repo := new(MockOrderRepository)

	service := OrderServiceImpl{
		OrderRepo: repo,
	}

	repo.On(
		"FindAllByUserID",
		mock.Anything,
		uint64(1),
	).Return(
		[]domain.Order{
			{ID: 1},
			{ID: 2},
		},
		nil,
	)

	result, err := service.FindAll(context.Background(), 1)
	assert.NoError(t, err)
	assert.Len(t, result, 2)

	repo.AssertExpectations(t)
}

func TestFindByIDOrderSuccess(t *testing.T) {
	repo := new(MockOrderRepository)

	service := OrderServiceImpl{
		OrderRepo: repo,
	}

	repo.On(
		"FindByIDAndUserID",
		mock.Anything,
		uint64(1),
		uint64(1),
	).Return(
		&domain.Order{
			ID:     1,
			UserID: 1,
		},
		nil,
	)

	result, err := service.FindByID(context.Background(), 1, 1)

	assert.NoError(t, err)
	assert.NotNil(t, result)

	repo.AssertExpectations(t)
}

func TestFindByIDOrderNotFound(t *testing.T) {
	repo := new(MockOrderRepository)

	service := OrderServiceImpl{
		OrderRepo: repo,
	}

	repo.On(
		"FindByIDAndUserID",
		mock.Anything,
		uint64(99),
		uint64(1),
	).Return(
		nil,
		gorm.ErrRecordNotFound,
	)

	result, err := service.FindByID(context.Background(), 99, 1)

	assert.Error(t, err)
	assert.Nil(t, result)

	repo.AssertExpectations(t)
}

func TestMarkAsShippedSuccess(t *testing.T) {
	repo := new(MockOrderRepository)

	service := OrderServiceImpl{
		OrderRepo: repo,
	}

	repo.On(
		"FindByID",
		mock.Anything,
		uint64(1),
	).Return(
		&domain.Order{
			ID:     1,
			Status: "PAID",
		},
		nil,
	)

	repo.On(
		"Update",
		mock.Anything,
		mock.AnythingOfType("*domain.Order"),
	).Return(nil)

	err := service.MarkAsShipped(context.Background(), 1)

	assert.NoError(t, err)

	repo.AssertExpectations(t)
}

func TestMarkAsShippedInvalidStatus(t *testing.T) {
	repo := new(MockOrderRepository)

	service := OrderServiceImpl{
		OrderRepo: repo,
	}

	repo.On(
		"FindByID",
		mock.Anything,
		uint64(1),
	).Return(
		&domain.Order{
			ID:     1,
			Status: "PENDING",
		},
		nil,
	)

	err := service.MarkAsShipped(context.Background(), 1)

	assert.Error(t, err)

	repo.AssertExpectations(t)
}

func TestMarkAsCompletedSuccess(t *testing.T) {
	repo := new(MockOrderRepository)

	service := OrderServiceImpl{
		OrderRepo: repo,
	}

	repo.On(
		"FindByID",
		mock.Anything,
		uint64(1),
	).Return(
		&domain.Order{
			ID:     1,
			Status: "SHIPPED",
		},
		nil,
	)

	repo.On(
		"Update",
		mock.Anything,
		mock.AnythingOfType("*domain.Order"),
	).Return(nil)

	err := service.MarkAsCompleted(context.Background(), 1)

	assert.NoError(t, err)

	repo.AssertExpectations(t)
}

func TestMarkAsCompletedInvalidStatus(t *testing.T) {
	repo := new(MockOrderRepository)

	service := OrderServiceImpl{
		OrderRepo: repo,
	}

	repo.On(
		"FindByID",
		mock.Anything,
		uint64(1),
	).Return(
		&domain.Order{
			ID:     1,
			Status: "PADI",
		},
		nil,
	)

	err := service.MarkAsCompleted(context.Background(), 1)

	assert.Error(t, err)

	repo.AssertExpectations(t)
}

func TestMarkAsShippedNotFound(t *testing.T) {
	repo := new(MockOrderRepository)

	service := OrderServiceImpl{
		OrderRepo: repo,
	}

	repo.On(
		"FindByID",
		mock.Anything,
		uint64(99),
	).Return(
		nil,
		gorm.ErrRecordNotFound,
	)

	err := service.MarkAsShipped(context.Background(), 99)

	assert.Error(t, err)

	repo.AssertExpectations(t)
}

func TestMarkAsCompletedNotFound(t *testing.T) {
	repo := new(MockOrderRepository)

	service := OrderServiceImpl{
		OrderRepo: repo,
	}

	repo.On(
		"FindByID",
		mock.Anything,
		uint64(99),
	).Return(
		nil,
		gorm.ErrRecordNotFound,
	)

	err := service.MarkAsCompleted(context.Background(), 99)

	assert.Error(t, err)

	repo.AssertExpectations(t)
}

