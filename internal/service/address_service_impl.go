package service

import (
	"context"

	"github.com/ardhisparahita/ecommerce-api/internal/domain"
	"github.com/ardhisparahita/ecommerce-api/internal/dto/request"
	"github.com/ardhisparahita/ecommerce-api/internal/dto/response"
	"github.com/ardhisparahita/ecommerce-api/internal/repository"
)

type AddressServiceImpl struct {
	Repo repository.AddressRepository
}

func NewAddressService(repo repository.AddressRepository) AddressService {
	return &AddressServiceImpl{
		Repo: repo,
	}
}

func (s *AddressServiceImpl) Create(ctx context.Context, userID uint64, req request.CreateAddressRequest) error {
	address := domain.Address{
		UserID:        userID,
		RecipientName: req.RecipientName,
		Phone:         req.Phone,
		Address:       req.Address,
		City:          req.City,
		PostalCode:    req.PostalCode,
	}

	return s.Repo.Create(ctx, &address)
}

func (s *AddressServiceImpl) FindAllByUserID(ctx context.Context, userID uint64) ([]response.AddressResponse, error) {
	addresses, err := s.Repo.FindAllByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	var result []response.AddressResponse

	for _, address := range addresses {
		result = append(result, response.AddressResponse{
			ID:            address.ID,
			RecipientName: address.RecipientName,
			Phone:         address.Phone,
			Address:       address.Address,
			City:          address.City,
			PostalCode:    address.PostalCode,
		})
	}

	return result, nil
}

func (s *AddressServiceImpl) FindByID(ctx context.Context, id uint64, userID uint64) (*response.AddressResponse, error) {
	address, err := s.Repo.FindByIDAndUserID(ctx, id, userID)
	if err != nil {
		return nil, err
	}

	return &response.AddressResponse{
		ID:            address.ID,
		RecipientName: address.RecipientName,
		Phone:         address.Phone,
		Address:       address.Address,
		City:          address.City,
		PostalCode:    address.PostalCode,
	}, nil
}

func (s *AddressServiceImpl) Update(ctx context.Context, id uint64, userID uint64, req request.UpdateAddressRequest) error {
	address, err := s.Repo.FindByIDAndUserID(ctx, id, userID)
	if err != nil {
		return err
	}

	address.RecipientName = req.RecipientName
	address.Phone = req.Phone
	address.Address = req.Address
	address.City = req.City
	address.PostalCode = req.PostalCode

	return s.Repo.Update(ctx, address)
}

func (s *AddressServiceImpl) Delete(ctx context.Context, id uint64, userID uint64) error {
	_, err := s.Repo.FindByIDAndUserID(ctx, id, userID)
	if err != nil {
		return err
	}

	return s.Repo.Delete(ctx, id)
}
