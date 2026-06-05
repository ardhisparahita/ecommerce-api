package service

import (
	"context"

	"github.com/ardhisparahita/ecommerce-api/internal/dto/request"
	"github.com/ardhisparahita/ecommerce-api/internal/dto/response"
)

type AuthService interface {
	Register(ctx context.Context, req request.RegisterRequest) error
	Login(ctx context.Context, req request.LoginRequest) (*response.AuthResponse, error)
}
