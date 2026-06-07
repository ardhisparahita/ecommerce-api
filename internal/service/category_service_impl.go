package service

import (
	"context"

	"github.com/ardhisparahita/ecommerce-api/internal/domain"
	"github.com/ardhisparahita/ecommerce-api/internal/dto/request"
	"github.com/ardhisparahita/ecommerce-api/internal/dto/response"
	"github.com/ardhisparahita/ecommerce-api/internal/mapper"
	"github.com/ardhisparahita/ecommerce-api/internal/repository"
)

type CategoryServiceImpl struct {
	Repo repository.CategoryRepository
}

func NewCategoryService(repo repository.CategoryRepository) CategoryService {
	return &CategoryServiceImpl{
		Repo: repo,
	}
}

func (s *CategoryServiceImpl) Create(ctx context.Context, req request.CreateCategoryRequest) (*response.CategoryResponse, error) {
	category := domain.Category{
		Name: req.Name,
	}
	err := s.Repo.Create(ctx, &category)
	if err != nil {
		return nil, err
	}

	return mapper.ToCategoryResponse(&category), nil
}

func (s *CategoryServiceImpl) FindAll(ctx context.Context) ([]response.CategoryResponse, error) {
	categories, err := s.Repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return mapper.ToCategoryResponses(categories), nil
}
