package service

import (
	"context"
	"math"

	"github.com/ardhisparahita/ecommerce-api/internal/domain"
	"github.com/ardhisparahita/ecommerce-api/internal/dto/request"
	"github.com/ardhisparahita/ecommerce-api/internal/dto/response"
	"github.com/ardhisparahita/ecommerce-api/internal/mapper"
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

func (s *ProductServiceImpl) Create(ctx context.Context, req request.CreateProductRequest) (*response.ProductResponse, error) {
	product := domain.Product{
		CategoryID:  req.CategoryID,
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
		ImageURL:    req.ImageURL,
	}

	err := s.Repo.Create(ctx, &product)
	if err != nil {
		return nil, err
	}

	return mapper.ToProductResponse(&product), nil

}

func (s *ProductServiceImpl) FindAll(ctx context.Context, req request.ProductQueryRequest) (*response.ProductListResponse, error) {
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.Limit <= 0 {
		req.Limit = 10
	}

	products, totalRows, err := s.Repo.FindAll(ctx, req)
	if err != nil {
		return nil, err
	}

	totalPage := int(math.Ceil(float64(totalRows) / float64(req.Limit)))

	return &response.ProductListResponse{
		Items:      mapper.ToProductResponses(products),
		Page:       req.Page,
		Limit:      req.Limit,
		TotalRows:  totalRows,
		TotalPages: totalPage,
	}, nil

}

func (s *ProductServiceImpl) FindByID(ctx context.Context, id uint64) (*response.ProductResponse, error) {
	product, err := s.Repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	res := &response.ProductResponse{
		ID:          product.ID,
		CategoryID:  product.CategoryID,
		Category:    product.Category.Name,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       product.Stock,
		ImageURL:    product.ImageURL,
	}

	return res, nil
}

func (s *ProductServiceImpl) Update(ctx context.Context, id uint64, req request.UpdateProductRequest) (*response.ProductResponse, error) {
	product, err := s.Repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	product.CategoryID = req.CategoryID
	product.Name = req.Name
	product.Description = req.Description
	product.Price = req.Price
	product.Stock = req.Stock
	product.ImageURL = req.ImageURL

	err = s.Repo.Update(ctx, product)
	if err != nil {
		return nil, err
	}

	res := &response.ProductResponse{
		ID:          product.ID,
		CategoryID:  product.CategoryID,
		Category:    product.Category.Name,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       product.Stock,
		ImageURL:    product.ImageURL,
	}

	return res, nil
}

func (s *ProductServiceImpl) Delete(ctx context.Context, id uint64) error {
	return s.Repo.Delete(ctx, id)
}
