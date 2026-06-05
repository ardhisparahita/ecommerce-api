package service

import (
	"context"

	"github.com/ardhisparahita/ecommerce-api/internal/domain"
	"github.com/ardhisparahita/ecommerce-api/internal/dto/request"
	"github.com/ardhisparahita/ecommerce-api/internal/dto/response"
	"github.com/ardhisparahita/ecommerce-api/internal/repository"
)

type ProductServiceImpl struct {
	Repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) ProductService {
	return &ProductServiceImpl{
		Repo: repo,
	}
}

func (s *ProductServiceImpl) Create(ctx context.Context, req request.CreateProductRequest) error {
	product := domain.Product{
		CategoryId:  req.CategoryID,
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
		ImageURL:    req.ImageURL,
	}

	return s.Repo.Create(ctx, &product)
}

func (s *ProductServiceImpl) FindAll(ctx context.Context) ([]response.ProductResponse, error) {
	products, err := s.Repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var result []response.ProductResponse

	for _, p := range products {
		result = append(result, response.ProductResponse{
			ID:          p.ID,
			CategoryID:  p.CategoryId,
			Category:    p.Category.Name,
			Name:        p.Name,
			Description: p.Description,
			Price:       p.Price,
			Stock:       p.Stock,
			ImageURL:    p.ImageURL,
		})
	}

	return result, nil

}

func (s *ProductServiceImpl) FindByID(ctx context.Context, id uint64) (*response.ProductResponse, error) {
	product, err := s.Repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	res := &response.ProductResponse{
		ID:          product.ID,
		CategoryID:  product.CategoryId,
		Category:    product.Category.Name,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       product.Stock,
		ImageURL:    product.ImageURL,
	}

	return res, nil
}

func (s *ProductServiceImpl) Update(ctx context.Context, id uint64, req request.UpdateProductRequest) error {
	product, err := s.Repo.FindByID(ctx, id)
	if err != nil {
		return nil
	}

	product.CategoryId = req.CategoryID
	product.Name = req.Name
	product.Description = req.Description
	product.Price = req.Price
	product.Stock = req.Stock
	product.ImageURL = req.ImageURL

	return s.Repo.Update(ctx, product)
}

func (s *ProductServiceImpl) Delete(ctx context.Context, id uint64) error {
	return s.Repo.Delete(ctx, id)
}
