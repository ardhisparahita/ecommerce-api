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

func TestMarkAsPaidSuccess(t *testing.T) {
	orderRepo := new(MockOrderRepository)
	paymentRepo := new(MockPaymentRepository)

	service := OrderServiceImpl{
		OrderRepo:   orderRepo,
		PaymentRepo: paymentRepo,
	}

	order := &domain.Order{
		ID:     1,
		Status: domain.OrderPending,
	}

	payment := &domain.Payment{
		OrderID: 1,
		Status:  domain.PaymentPending,
	}

	orderRepo.On(
		"FindByID",
		mock.Anything,
		uint64(1),
	).Return(order, nil)

	paymentRepo.On(
		"FIndByOrderID",
		mock.Anything,
		uint64(1),
	).Return(payment, nil)

	paymentRepo.On(
		"UpdateTx",
		mock.Anything,
		mock.Anything,
		payment,
	).Return(nil)

	orderRepo.On(
		"UpdateTx",
		mock.Anything,
		mock.Anything,
		order,
	).Return(nil)

	err := service.MarkAsPaid(context.Background(), 1)

	assert.NoError(t, err)
	assert.Equal(t, domain.OrderPaid, order.Status)
	assert.Equal(t, domain.PaymentPaid, payment.Status)

	orderRepo.AssertExpectations(t)
	paymentRepo.AssertExpectations(t)
}

func TestMarkAsPaidOrderNotFound(t *testing.T) {

	repo := new(MockOrderRepository)

	service := &OrderServiceImpl{
		OrderRepo: repo,
	}

	repo.
		On("FindByID", mock.Anything, uint64(99)).
		Return(nil, gorm.ErrRecordNotFound)

	err := service.MarkAsPaid(context.Background(), 99)

	assert.Error(t, err)

	repo.AssertExpectations(t)
}

func TestMarkAsPaidPaymentNotFound(t *testing.T) {

	orderRepo := new(MockOrderRepository)
	paymentRepo := new(MockPaymentRepository)

	service := &OrderServiceImpl{
		OrderRepo:   orderRepo,
		PaymentRepo: paymentRepo,
	}

	order := &domain.Order{
		ID:     1,
		Status: domain.OrderPending,
	}

	orderRepo.
		On("FindByID", mock.Anything, uint64(1)).
		Return(order, nil)

	paymentRepo.
		On("FIndByOrderID", mock.Anything, uint64(1)).
		Return(nil, gorm.ErrRecordNotFound)

	err := service.MarkAsPaid(context.Background(), 1)

	assert.Error(t, err)

	orderRepo.AssertExpectations(t)
	paymentRepo.AssertExpectations(t)
}

func TestMarkAsPaidPaymentNotFound(t *testing.T) {

	orderRepo := new(MockOrderRepository)
	paymentRepo := new(MockPaymentRepository)

	service := &OrderServiceImpl{
		OrderRepo:   orderRepo,
		PaymentRepo: paymentRepo,
	}

	order := &domain.Order{
		ID:     1,
		Status: domain.OrderPending,
	}

	orderRepo.
		On("FindByID", mock.Anything, uint64(1)).
		Return(order, nil)

	paymentRepo.
		On("FIndByOrderID", mock.Anything, uint64(1)).
		Return(nil, gorm.ErrRecordNotFound)

	err := service.MarkAsPaid(context.Background(), 1)

	assert.Error(t, err)

	orderRepo.AssertExpectations(t)
	paymentRepo.AssertExpectations(t)
}

func TestMarkAsPaidAlreadyPaid(t *testing.T) {

	repo := new(MockOrderRepository)

	service := &OrderServiceImpl{
		OrderRepo: repo,
	}

	order := &domain.Order{
		Status: domain.OrderPaid,
	}

	repo.
		On("FindByID", mock.Anything, uint64(1)).
		Return(order, nil)

	err := service.MarkAsPaid(context.Background(), 1)

	assert.Error(t, err)

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
			Status: "PAID",
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
