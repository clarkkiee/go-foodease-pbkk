package service

import (
	"context"
	"fmt"
	"go-foodease-be/dto"
	"go-foodease-be/models"
	"go-foodease-be/repository"

	"github.com/google/uuid"
	"github.com/leodido/go-encodeuricomponent"
)

type (
	AddressService interface {
		CreateNewAddress(ctx context.Context, req dto.CreateNewAddressRequest, id string) (dto.AddressResponse, error)
		GetAllAddressByCustomerId(ctx context.Context, id string) ([]dto.AddressResponse, error)
		GetAddressById(ctx context.Context, addressId string, customerId string) (dto.AddressResponse, error)
	}

	addressService struct {
		addressRepo repository.AddressRepository
		jwtService JWTService
	}
)

func NewAddressService(addressRepo repository.AddressRepository, jwtService JWTService) AddressService {
	return &addressService{
		addressRepo: addressRepo,
		jwtService: jwtService,
	}
}

func (s *addressService) CreateNewAddress(ctx context.Context, req dto.CreateNewAddressRequest, id string) (dto.AddressResponse, error) {
	fullAddr := fmt.Sprintf("%s, %s, %s, %s, %s", req.Street, req.Village, req.SubDistrict, req.City, req.Province)
	encoded := encodeuricomponent.EncodeURIComponent(fullAddr)
	
	coords, err := s.addressRepo.ProduceCordFromText(ctx, nil, encoded)
	if err != nil {
		return dto.AddressResponse{}, err
	}

	newAddress := models.Address{
		Street: fullAddr,
		Coordinates: *coords,
		CustomerID: uuid.MustParse(id),
	}

	res, err := s.addressRepo.CreateAddress(ctx, nil, newAddress)
	if err != nil {
		return dto.AddressResponse{}, err
	}

	return dto.AddressResponse{
		ID: res.ID,
		Street: res.Street,
		Longitude: res.Coordinates.Longitude,
		Latitude: res.Coordinates.Latitude,
		CreatedAt: res.CreatedAt,
		UpdatedAt: res.UpdatedAt,
	}, nil

}

func (s *addressService) GetAllAddressByCustomerId(ctx context.Context, customerId string) ([]dto.AddressResponse, error){
	addresses, err := s.addressRepo.GetAllAddressByCustomerId(ctx, nil, customerId)	
	if err != nil {
		return []dto.AddressResponse{}, err
	}

	return addresses, nil
}

func (s *addressService) GetAddressById(ctx context.Context, addressId string, customerId string) (dto.AddressResponse, error){
	addr, err := s.addressRepo.GetAddressById(ctx, nil, addressId, customerId)
	if err != nil {
		return dto.AddressResponse{}, err
	}

	return addr, nil
}