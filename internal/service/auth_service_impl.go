package service

import (
	"context"

	"github.com/ardhisparahita/ecommerce-api/internal/domain"
	"github.com/ardhisparahita/ecommerce-api/internal/dto/request"
	"github.com/ardhisparahita/ecommerce-api/internal/dto/response"
	"github.com/ardhisparahita/ecommerce-api/internal/repository"
	"github.com/ardhisparahita/ecommerce-api/pkg/config"
	"github.com/ardhisparahita/ecommerce-api/pkg/utils"
	"golang.org/x/crypto/bcrypt"
)

type AuthServiceImpl struct {
	Repo repository.UserRepository
}

func NewAuthService(repo repository.UserRepository) AuthService {
	return &AuthServiceImpl{
		Repo: repo,
	}
}

func (s *AuthServiceImpl) Register(ctx context.Context, req request.RegisterRequest) error {
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return err
	}

	user := domain.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashedPassword),
	}

	return s.Repo.Create(ctx, &user)
}

func (s *AuthServiceImpl) Login(ctx context.Context, req request.LoginRequest) (*response.AuthResponse, error) {
	user, err := s.Repo.FindByEmail(ctx, req.Email)
	if err != nil {
		return nil, utils.Unauthorized("invalid email or password")
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(req.Password),
	)
	if err != nil {
		return nil, utils.Unauthorized("invalid email or password")
	}

	token, err := utils.GenerateToken(
		user.ID,
		config.Get("JWT_SECRET"),
	)
	if err != nil {
		return nil, err
	}

	return &response.AuthResponse{Token: token}, nil
}
