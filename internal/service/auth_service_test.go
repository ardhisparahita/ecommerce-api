package service

import (
	"context"
	"testing"

	"github.com/ardhisparahita/ecommerce-api/internal/domain"
	"github.com/ardhisparahita/ecommerce-api/internal/dto/request"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Create(
	ctx context.Context,
	user *domain.User,
) error {
	args := m.Called(ctx, user)

	return args.Error(0)
}

func (m *MockUserRepository) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	args := m.Called(ctx, email)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*domain.User), args.Error(1)
}

func (m *MockUserRepository) FindByID(ctx context.Context, id uint64) (*domain.User, error) {
	args := m.Called(ctx, id)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*domain.User), args.Error(1)
}

func (m *MockUserRepository) Update(ctx context.Context, user *domain.User) error {
	args := m.Called(ctx, user)

	return args.Error(0)
}

func TestRegisterSuccess(t *testing.T) {
	repo := new(MockUserRepository)

	service := AuthServiceImpl{
		Repo: repo,
	}

	req := request.RegisterRequest{
		Name:     "Ardhis",
		Email:    "ardhis@gmail.com",
		Password: "password123",
	}

	repo.On("FindByEmail", mock.Anything, req.Email).Return(nil, gorm.ErrRecordNotFound)
	repo.On("Create", mock.Anything, mock.AnythingOfType("*domain.User")).Return(nil)

	result, err := service.Register(context.Background(), req)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, req.Email, result.Email)

	repo.AssertExpectations(t)
}

func TestRegisterDuplicateEmail(t *testing.T) {
	repo := new(MockUserRepository)

	service := AuthServiceImpl{
		Repo: repo,
	}

	repo.On("FindByEmail", mock.Anything, "ardhis@gmail.com").Return(&domain.User{
		ID:    1,
		Email: "ardhis@gmail.com",
	}, nil)

	result, err := service.Register(context.Background(), request.RegisterRequest{
		Name:     "Ardhis",
		Email:    "ardhis@gmail.com",
		Password: "password123",
	})

	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestLoginSuccess(t *testing.T) {
	hashedPassword, _ := bcrypt.GenerateFromPassword(
		[]byte("password123"),
		bcrypt.DefaultCost,
	)

	repo := new(MockUserRepository)

	service := AuthServiceImpl{
		Repo: repo,
	}

	repo.On(
		"FindByEmail",
		mock.Anything,
		"ardhis@gmail.com",
	).Return(
		&domain.User{
			ID:       1,
			Email:    "ardhis@gmail.com",
			Password: string(hashedPassword),
			Role:     "CUSTOMER",
		},
		nil,
	)

	result, err := service.Login(
		context.Background(),
		request.LoginRequest{
			Email:    "ardhis@gmail.com",
			Password: "password123",
		},
	)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.NotEmpty(t, result.Token)
}

func TestLoginWrongPassword(t *testing.T) {
	hashedPassword, _ := bcrypt.GenerateFromPassword(
		[]byte("password123"),
		bcrypt.DefaultCost,
	)

	repo := new(MockUserRepository)

	service := AuthServiceImpl{
		Repo: repo,
	}

	repo.On(
		"FindByEmail",
		mock.Anything,
		"ardhis@gmail.com",
	).Return(
		&domain.User{
			ID:       1,
			Email:    "ardhis@gmail.com",
			Password: string(hashedPassword),
			Role:     "CUSTOMER",
		},
		nil,
	)

	result, err := service.Login(
		context.Background(),
		request.LoginRequest{
			Email:    "ardhis@gmail.com",
			Password: "salahpassword",
		},
	)

	assert.Error(t, err)
	assert.Nil(t, result)
}
