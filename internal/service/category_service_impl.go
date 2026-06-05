package service

import (
	"context"

	"github.com/ardhisparahita/ecommerce-api/internal/domain"
	"github.com/ardhisparahita/ecommerce-api/internal/dto/request"
	"github.com/ardhisparahita/ecommerce-api/internal/dto/response"
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

	res := &response.CategoryResponse{
		ID:   category.ID,
		Name: category.Name,
	}

	return res, nil
}

func (s *CategoryServiceImpl) FindAll(ctx context.Context) ([]response.CategoryResponse, error) {
	categories, err := s.Repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var result []response.CategoryResponse

	for _, c := range categories {
		result = append(result, response.CategoryResponse{
			ID:   c.ID,
			Name: c.Name,
		})
	}

	return result, nil
}
