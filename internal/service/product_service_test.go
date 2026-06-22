package service

import (
	"context"
	"testing"

	"github.com/ardhisparahita/ecommerce-api/internal/domain"
	"github.com/ardhisparahita/ecommerce-api/internal/dto/request"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

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

	return args.Get(0).([]domain.Product), args.Get(1).(int64), args.Error(2)
}

func (m *MockProductRepository) FindByID(ctx context.Context, id uint64) (*domain.Product, error) {
	args := m.Called(ctx, id)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*domain.Product), args.Error(1)
}

func (m *MockProductRepository) Update(ctx context.Context, product *domain.Product) error {
	args := m.Called(ctx, product)

	return args.Error(0)
}

func (m *MockProductRepository) UpdateTx(ctx context.Context, tx *gorm.DB, product *domain.Product) error {
	args := m.Called(ctx, tx, product)

	return args.Error(0)
}

func (m *MockProductRepository) Delete(ctx context.Context, id uint64) error {
	args := m.Called(ctx, id)

	return args.Error(0)
}

func TestCreateProductSuccess(t *testing.T) {
	repo := new(MockProductRepository)

	service := ProductServiceImpl{
		Repo: repo,
	}

	req := request.CreateProductRequest{
		CategoryID:  1,
		Name:        "Keyboard",
		Description: "Mechanical",
		Price:       100000,
		Stock:       10,
	}

	repo.On(
		"Create",
		mock.Anything,
		mock.AnythingOfType("*domain.Product"),
	).Run(func(args mock.Arguments) {
		product := args.Get(1).(*domain.Product)

		product.ID = 1
	}).Return(nil)

	repo.On(
		"FindByID",
		mock.Anything,
		uint64(1),
	).Return(
		&domain.Product{
			ID:         1,
			CategoryID: 1,
			Name:       "Keyboard",
		},
		nil,
	)

	result, err := service.Create(
		context.Background(),
		req,
	)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "Keyboard", result.Name)

	repo.AssertExpectations(t)
}

func TestFindByIDSuccess(t *testing.T) {
	repo := new(MockProductRepository)

	service := ProductServiceImpl{
		Repo: repo,
	}

	repo.On(
		"FindByID",
		mock.Anything,
		uint64(1),
	).Return(
		&domain.Product{
			ID:   1,
			Name: "Keyboard",
		},
		nil,
	)

	result, err := service.FindByID(context.Background(), 1)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, uint64(1), result.ID)

	repo.AssertExpectations(t)
}

func TestFindByIDNotFound(t *testing.T) {
	repo := new(MockProductRepository)

	service := ProductServiceImpl{
		Repo: repo,
	}

	repo.On(
		"FindByID",
		mock.Anything,
		uint64(99),
	).Return(
		nil,
		gorm.ErrRecordNotFound,
	)

	result, err := service.FindByID(context.Background(), 99)

	assert.Error(t, err)
	assert.Nil(t, result)

	repo.AssertExpectations(t)
}

func TestDeleteProductSuccess(t *testing.T) {
	repo := new(MockProductRepository)

	service := ProductServiceImpl{
		Repo: repo,
	}

	repo.On(
		"FindByID",
		mock.Anything,
		uint64(1),
	).Return(
		&domain.Product{
			ID: 1,
		},
		nil,
	)

	repo.On(
		"Delete",
		mock.Anything,
		uint64(1),
	).Return(nil)

	err := service.Delete(context.Background(), 1)

	assert.NoError(t, err)

	repo.AssertExpectations(t)
}

func TestUploadImageSuccess(t *testing.T) {
	repo := new(MockProductRepository)

	service := ProductServiceImpl{
		Repo: repo,
	}

	repo.On(
		"FindByID",
		mock.Anything,
		uint64(1),
	).Return(
		&domain.Product{
			ID:   1,
			Name: "Keyboard",
		},
		nil,
	)

	repo.On(
		"Update",
		mock.Anything,
		mock.AnythingOfType("*domain.Product"),
	).Return(nil)

	result, err := service.UploadImage(context.Background(), 1, "/uploads/products/test.jpg")

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "/uploads/products/test.jpg", result.ImageURL)

	repo.AssertExpectations(t)
}

func TestUpdateProductSuccess(t *testing.T) {
	repo := new(MockProductRepository)

	service := ProductServiceImpl{
		Repo: repo,
	}

	repo.On(
		"FindByID",
		mock.Anything,
		uint64(1),
	).Return(
		&domain.Product{
			ID:   1,
			Name: "Old Product",
		},
		nil,
	)

	repo.On(
		"Update",
		mock.Anything,
		mock.Anything,
	).Return(nil)

	result, err := service.Update(context.Background(), 1, request.UpdateProductRequest{
		CategoryID:  1,
		Name:        "New Product",
		Description: "Updated",
		Price:       100000,
		Stock:       20,
	})

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "New Product", result.Name)

	repo.AssertExpectations(t)
}

func TestUpdateProductNotFound(t *testing.T) {
	repo := new(MockProductRepository)

	service := ProductServiceImpl{
		Repo: repo,
	}

	repo.On(
		"FindByID",
		mock.Anything,
		uint64(99),
	).Return(
		nil,
		gorm.ErrRecordNotFound,
	)

	result, err := service.Update(context.Background(), 99, request.UpdateProductRequest{})

	assert.Error(t, err)
	assert.Nil(t, result)

	repo.AssertExpectations(t)
}

func TestDeleteProductNotFo(t *testing.T) {
	repo := new(MockProductRepository)

	service := ProductServiceImpl{
		Repo: repo,
	}

	repo.On(
		"FindByID",
		mock.Anything,
		uint64(99),
	).Return(
		nil,
		gorm.ErrRecordNotFound,
	)

	err := service.Delete(
		context.Background(),
		99,
	)

	assert.Error(t, err)

	repo.AssertExpectations(t)
}

func TestUploadImageNotFound(t *testing.T) {
	repo := new(MockProductRepository)

	service := ProductServiceImpl{
		Repo: repo,
	}

	repo.On(
		"FindByID",
		mock.Anything,
		uint64(99),
	).Return(
		nil,
		gorm.ErrRecordNotFound,
	)

	result, err := service.UploadImage(context.Background(), 99, "test.jpg")
	assert.Error(t, err)
	assert.Nil(t, result)

	repo.AssertExpectations(t)
}

func TestFindAllSuccess(t *testing.T) {
	repo := new(MockProductRepository)

	service := ProductServiceImpl{
		Repo: repo,
	}

	req := request.ProductQueryRequest{
		Page:  1,
		Limit: 10,
	}

	repo.On(
		"FindAll",
		mock.Anything,
		req,
	).Return(
		[]domain.Product{
			{
				ID:   1,
				Name: "Keyboard",
			},
			{
				ID:   2,
				Name: "Mouse",
			},
		},
		int64(2),
		nil,
	)

	result, err := service.FindAll(context.Background(), req)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Len(t, result.Items, 2)

	repo.AssertExpectations(t)
}
